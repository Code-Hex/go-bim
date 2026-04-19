#!/usr/bin/env bash

set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
out_dir="${root_dir}/ifc/reference/express"

mkdir -p "${out_dir}"

curl -fsSL \
  "https://standards.buildingsmart.org/IFC/RELEASE/IFC2x3/TC1/EXPRESS/IFC2X3_TC1.exp" \
  -o "${out_dir}/IFC2X3_TC1.exp"

curl -fsSL \
  "https://standards.buildingsmart.org/IFC/RELEASE/IFC4/ADD2_TC1/EXPRESS/IFC4.exp" \
  -o "${out_dir}/IFC4_ADD2_TC1.exp"

curl -fsSL \
  "https://standards.buildingsmart.org/IFC/RELEASE/IFC4_3/HTML/IFC4X3_ADD2.exp" \
  -o "${out_dir}/IFC4X3_ADD2.exp"

shasum -a 256 \
  "${out_dir}/IFC2X3_TC1.exp" \
  "${out_dir}/IFC4_ADD2_TC1.exp" \
  "${out_dir}/IFC4X3_ADD2.exp"
