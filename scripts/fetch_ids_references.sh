#!/usr/bin/env bash

set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
official_dir="${root_dir}/ids/reference/official"
out_dir="${official_dir}/v1.0.0"
tmp_out_dir="${official_dir}/.v1.0.0.tmp"
tmp_sources="${official_dir}/.SOURCES.txt.tmp"
snapshot_date="$(date +%F)"

rm -rf "${tmp_out_dir}"
mkdir -p \
  "${tmp_out_dir}/Schema" \
  "${tmp_out_dir}/Documentation" \
  "${tmp_out_dir}/Documentation/testcases/ids" \
  "${tmp_out_dir}/Documentation/testcases/entity" \
  "${tmp_out_dir}/Documentation/testcases/attribute" \
  "${tmp_out_dir}/Documentation/testcases/classification" \
  "${tmp_out_dir}/Documentation/testcases/property" \
  "${tmp_out_dir}/Documentation/testcases/material" \
  "${tmp_out_dir}/Documentation/testcases/partof" \
  "${tmp_out_dir}/Documentation/testcases/restriction" \
  "${tmp_out_dir}/Documentation/testcases/tolerance"

curl -fsSL \
  "https://standards.buildingsmart.org/IDS/1.0/ids.xsd" \
  -o "${tmp_out_dir}/Schema/ids.xsd"

docs=(
  "Documentation/README.md"
  "Documentation/DataTypes.md"
  "Documentation/attribute-facet.md"
  "Documentation/classification-facet.md"
  "Documentation/developer-guide.md"
  "Documentation/entity-facet.md"
  "Documentation/facet-configurations.md"
  "Documentation/ids-metadata.md"
  "Documentation/ifc2x3-occurrence-type-mapping-table.md"
  "Documentation/material-facet.md"
  "Documentation/partof-facet.md"
  "Documentation/property-facet.md"
  "Documentation/restrictions.md"
  "Documentation/specifications.md"
  "Documentation/testcases/contributing.md"
  "Documentation/testcases/scripts.md"
)

examples=(
  "Documentation/testcases/ids/pass-a_minimal_ids_can_check_a_minimal_ifc_2_2.ids"
  "Documentation/testcases/entity/pass-a_matching_predefined_type_should_pass.ids"
  "Documentation/testcases/attribute/pass-typecast_checking_may_also_occur_within_enumeration_restrictions.ids"
  "Documentation/testcases/classification/pass-restrictions_can_be_used_for_values_2_3.ids"
  "Documentation/testcases/property/pass-measures_are_used_to_specify_an_ifc_data_type_2_2.ids"
  "Documentation/testcases/material/pass-any_material_name_in_a_list_will_pass_a_value_check.ids"
  "Documentation/testcases/partof/pass-an_aggregate_may_specify_the_entity_of_the_whole_1_2.ids"
  "Documentation/testcases/restriction/pass-max_and_min_length_checks_can_be_used_3_3.ids"
  "Documentation/testcases/tolerance/pass-comparison_tolerance_for_floating_point_positive_high_number_lower_bound.ids"
)

for path in "${docs[@]}" "${examples[@]}"; do
  destination="${tmp_out_dir}/${path}"
  mkdir -p "$(dirname "${destination}")"
  curl -fsSL \
    "https://raw.githubusercontent.com/buildingSMART/IDS/v1.0.0/${path}" \
    -o "${destination}"
done

(
  cd "${tmp_out_dir}"
  find . -type f ! -name 'SHA256SUMS' -print0 | sort -z | xargs -0 shasum -a 256 > SHA256SUMS
)

cat > "${tmp_sources}" <<EOF
buildingSMART IDS reference sources for go-bim

Snapshot date: ${snapshot_date}

Official IDS standards confirmed from primary sources:
- v1.0.0 is the only official buildingSMART IDS standard.
  - GitHub release: https://github.com/buildingSMART/IDS/releases/tag/v1.0.0
    published 2024-06-03T15:52:52Z
  - buildingSMART page: https://www.buildingsmart.org/standards/bsi-standards/information-delivery-specification-ids/
    states IDS v1.0 is the final standard and previous versions were beta iterations.

Non-official but relevant compatibility reference:
- 0.9.7 candidate schema: https://standards.buildingsmart.org/IDS/0.9.7/ids.xsd

Vendored under ids/reference/official/v1.0.0:
- Schema/ids.xsd downloaded from https://standards.buildingsmart.org/IDS/1.0/ids.xsd
- Documentation/*.md plus Documentation/testcases/*.md downloaded from https://github.com/buildingSMART/IDS/tree/v1.0.0/Documentation
- Representative official .ids examples downloaded from https://github.com/buildingSMART/IDS/tree/v1.0.0/Documentation/testcases
 - SHA256SUMS records checksums for every vendored file in v1.0.0
EOF

rm -rf "${out_dir}"
mv "${tmp_out_dir}" "${out_dir}"
mv "${tmp_sources}" "${official_dir}/SOURCES.txt"

shasum -a 256 "${out_dir}/Schema/ids.xsd"
