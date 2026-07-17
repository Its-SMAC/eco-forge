package manifest

type ManifestTemplate string

const ForgeToml ManifestTemplate = `[Project]
name = "%s"
framework = "%s"
description = "%s"

ingots = [%s]

[Forge]
version = "%s"
`
