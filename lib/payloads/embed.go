package payloads

import (
	"embed"
)

//go:embed behinder/java/*.class
var BeClassFiles embed.FS

//go:embed behinder/php/*.php
var BePhpFiles embed.FS

//go:embed behinder/csharp/*.dll
var BeDllFiles embed.FS

//go:embed behinder/asp/*.asp
var BeAspFiles embed.FS

//go:embed godzilla/java/payloadv4.class
var GodClassFiles []byte

//go:embed godzilla/java/plugins/*.class
var GodClassPluginsFiles embed.FS

//go:embed godzilla/php/payloadv4.php
var GodPhpFiles []byte

//go:embed godzilla/php/plugins/*.php
var GodPhpPluginsFiles embed.FS

//go:embed godzilla/csharp/payload.dll
var GodDllFiles []byte

//go:embed godzilla/csharp/plugins/*.dll
var GodDllPluginsFiles embed.FS

//go:embed godzilla/asp/payload.asp
var GodAspFiles []byte

//go:embed godzilla/asp/plugins/*.asp
var GodAspPluginsFiles embed.FS
