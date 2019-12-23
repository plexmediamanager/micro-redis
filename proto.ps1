$goPathFromEnvironment = Get-Content env:\GOPATH
$goPathSources = Join-Path $goPathFromEnvironment "src"
$currentDirectory = $PSScriptRoot
$protobufFiles = Get-ChildItem $currentDirectory -File | Where-Object { $_.Extension -eq ".proto" }

foreach ($proto in $protobufFiles) {
    protoc -I $goPathSources --proto_path=. --micro_out=proto --go_out=proto $proto
}