# Test data

One `<Rule>.ct` file per fixture, holding the exact Vale output for that
fixture — the record of what the rule does today.

Generate and regenerate these with `make update` (`go test ./... -update`);
never edit them by hand. The diff they produce is how rule changes get reviewed.
