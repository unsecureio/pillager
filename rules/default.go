package rules

// DefaultConfig is the default ruleset for pillager's hunting parameters.
// This can be overridden by providing a rules.toml file as an argument.
const DefaultConfig = `
title = "pillager config"
[[rules]]
	description = "AWS Access Key"
	regex = '''(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}'''
	tags = ["key", "AWS"]
[[rules]]
	description = "AWS Secret Key"
	regex = '''(?i)aws(.{0,20})?(?-i)['\"][0-9a-zA-Z\/+]{40}['\"]'''
	tags = ["key", "AWS"]
[[rules]]
	description = "Github"
	regex = '''(?i)github(.{0,20})?(?-i)[0-9a-zA-Z]{35,40}'''
	tags = ["key", "Github"]
[[rules]]
	description = "Slack"
	regex = '''xox[baprs]-([0-9a-zA-Z]{10,48})?'''
	tags = ["key", "Slack"]
[[rules]]
	description = "Asymmetric Private Key"
	regex = '''-----BEGIN ((EC|PGP|DSA|RSA|OPENSSH) )?PRIVATE KEY( BLOCK)?-----'''
	tags = ["key", "AsymmetricPrivateKey"]
[[rules]]
	description = "Slack Webhook"
	regex = '''https://hooks.slack.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}'''
	tags = ["key", "slack"]

[allowlist]
	description = "Allowlisted files"
	files = ['''^\.?gitleaks.toml$''',
	'''(.*?)(png|jpg|gif|doc|docx|pdf|bin|xls|pyc|zip)$''',
	'''(go.mod|go.sum)$''']
`
