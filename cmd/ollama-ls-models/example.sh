./ollama-ls-models |
	jq -c |
	fgrep phi4:14b |
	jq '.license = null' |
	jq '.tensors = null' |
	jq '.modelfile = null' |
	jq '.template = null' |
	jq '.parameters=null' |
	dasel --read=json --write=toml | bat --language=toml
