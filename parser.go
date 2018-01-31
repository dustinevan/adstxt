package adstxt

func parse(txt string) ([]Path, []Variable, error) {
	//used to ensure paths are unique and that if a an acct is used for both reseller and direct it's marked as both
	pathmap := make(map[string]int)

	lines := strings.Split(txt, "\n")
	paths := make([]Path, 0)
	vars := make([]Variable, 0)

	// lines may be blank or not paths, this keeps track of which pathidx we're on. This index is
	// used in the duplicate and DIRECT+RESELLER -> BOTH check
	pathidx := 0
	// TODO: this is a long for loop, refactor
	for _, line := range lines {
		if line == "" {
			continue
		}
		// if line is a variable parse and continue
		if strings.Contains(line, "=") {
			part := strings.Split(line, "=")
			v := Variable{Key: part[0], Value: part[1]}
			vars = append(vars, v)
			continue
		}

		path := Path{}

		// remove whitespace
		line = strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				// if the character is a space, drop it
				return -1
			}
			// else keep it in the string
			return r
		}, line)

		// strip comments
		if strings.Contains(line, "#") {
			line = strings.Split(line, "#")[0]
		}

		// get extensions
		if strings.Contains(line, ";") {
			s := strings.Split(line, ";")
			line = s[0]
			path.Ext = s[1:]
		}

		fields := strings.Split(line, ",")
		if len(fields) < 3 {
			//log.Printf("line: |%s| does not contain all required fields, skipping.", line)
			continue
		}

		if !strings.Contains(fields[0], ".") {
			continue
		}

		// there are lots of ways of saying the same thing right now. This transfers
		// exchange domains to what they should be.
		exdom := strings.ToLower(fields[0])
		if name, ok := exdoms[exdom]; ok {
			if canonical, ok := exnames_canonical[name]; ok {
				exdom = canonical
			}
		}
		path.ExchangeDomain = exdom

		if fields[1] == enum.DIRECT.String() || fields[1] == enum.RESELLER.String() {
			continue
		}
		path.PubID = fields[1]

		at := GetAccountType(fields[2])
		if at == enum.INVALID_ACCT_TYPE {
			continue
		}
		path.Acct = at

		// TODO: put this in a function
		// check for duplicate paths, or duplicate paths that have DIRECT and RESELLER
		if idx, ok := pathmap[path.ExchangeDomain+path.PubID]; ok {
			if paths[idx].Acct == enum.DIRECT && path.Acct == enum.RESELLER ||
				paths[idx].Acct == enum.RESELLER && path.Acct == enum.DIRECT {
				paths[idx].Acct = enum.BOTH
			}
			continue
		}
		pathmap[path.ExchangeDomain+path.PubID] = pathidx
		// end duplicate checking

		if len(fields) > 3 {
			path.TrustID = fields[3]
		}
		paths = append(paths, path)
		pathidx++
	}
	return paths, vars, nil
}
