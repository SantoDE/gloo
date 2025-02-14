
---
title: "domain.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `xds.type.matcher.v3` 
#### Types:


- [ServerNameMatcher](#servernamematcher)
- [DomainMatcher](#domainmatcher)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/external/xds/type/matcher/v3/domain.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/external/xds/type/matcher/v3/domain.proto)





---
### ServerNameMatcher

 
Matches a fully qualified server name against a set of domain
names with optional wildcards.

```yaml
"domainMatchers": []xds.type.matcher.v3.ServerNameMatcher.DomainMatcher

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `domainMatchers` | [[]xds.type.matcher.v3.ServerNameMatcher.DomainMatcher](../domain.proto.sk/#domainmatcher) | Match a server name by multiple domain matchers. Each domain, exact or wildcard, must appear at most once across all the domain matchers. The server name will be matched against all wildcard domains starting from the longest suffix, i.e. ``www.example.com`` input will be first matched against ``www.example.com``, then ``*.example.com``, then ``*.com``, then ``*``, until the associated matcher action accepts the input. Note that wildcards must be on a dot border, and values like ``*w.example.com`` are invalid. |




---
### DomainMatcher

 
Specifies a set of exact and wildcard domains and a match action. The
wildcard symbol ``*`` must appear at most once as the left-most part of
the domain on a dot border. The wildcard matches one or more non-empty
domain parts.

```yaml
"domains": []string
"onMatch": .xds.type.matcher.v3.Matcher.OnMatch

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `domains` | `[]string` | A non-empty set of domain names with optional wildcards, e.g. ``www.example.com``, ``*.com``, or ``*``. |
| `onMatch` | [.xds.type.matcher.v3.Matcher.OnMatch](../matcher.proto.sk/#onmatch) | Match action to apply when the server name matches any of the domain names in the matcher. |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
