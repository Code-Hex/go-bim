# Third-Party Notices

This repository contains original Go source code and vendored reference
materials under different licenses.

## buildingSMART IDS official reference material

- Location: `ids/reference/official/**`
- Source repository: <https://github.com/buildingSMART/IDS>
- Source release used for vendored reference files: `v1.0.0`
- Upstream license notice: <https://github.com/buildingSMART/IDS/blob/development/LICENSE>
- License: Creative Commons Attribution-NoDerivatives 4.0 International
  (`CC BY-ND 4.0`)
- Canonical license URL: <https://creativecommons.org/licenses/by-nd/4.0/>

The vendored files under `ids/reference/official/**` are distributed as copied
reference material and are excluded from the repository's MIT license. Keep
them attributable to buildingSMART International Ltd. and do not relicense
them as MIT.

## buildingSMART IFC official EXPRESS schemas

- Location: `ifc/reference/express/**`
- Source site: <https://technical.buildingsmart.org/standards/ifc/ifc-schema-specifications/>
- Upstream schema URLs are recorded in `ifc/reference/express/SOURCES.txt`
- License: Creative Commons Attribution-NoDerivatives 4.0 International
  (`CC BY-ND 4.0`)
- Canonical license URL: <https://creativecommons.org/licenses/by-nd/4.0/>

The vendored files under `ifc/reference/express/**` are distributed as copied
official IFC schema material and are excluded from the repository's MIT
license. Keep them attributable to buildingSMART International Ltd. and do not
relicense them as MIT.

## ST-Bridge reference material and samples

- Locations: `stbridge/reference/**`, `stbridge/testdata/samples/**`
- Public source site used during import review:
  <https://www.building-smart.or.jp/meeting/buildall/structural-design/>
- Local provenance notes: `stbridge/reference/SOURCES.txt`,
  `stbridge/reference/sources/*.md`

These files were imported with the `stbridge` package from the original
`go-stbridge` worktree and include ST-Bridge XSD files, transcription notes,
and sample documents. No explicit public reuse license was identified during
import review, so they are excluded from this repository's MIT license. This
repository does not relicense them as MIT.
