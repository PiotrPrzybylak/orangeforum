package templates

const profileSrc = `
{{ define "head" }}
<style>
.row {
	margin-top: 10px;
}
input[type="text"], input[type="number"], input[type="email"], textarea {
	width: 90%;
}
@media screen and (min-width:600px) {
	.col1 {
		float: left;
		text-align: right;
		width: 150px;
	}
	.col2 {
		float: left;
		text-align: left;
		margin-left: 15px;
		width: 300px;
	}
	.col1-offset {
		margin-left: 165px;
	}
}
</style>
{{ end }}


{{ define "content" }}


<form action="/users" method="POST">
<input type="hidden" name="csrf" value="{{ .Common.CSRF }}">
<div class="row clearfix">
	<div class="col1">Username</div>
	<div class="col2">{{ .UserName }}</div>
</div>
<div class="row clearfix">
	<div class="col1">Karma</div>
	<div class="col2">{{ .Karma }}</div>
</div>
<div class="row clearfix">
	<div class="col1">About (public)</div>
	<div class="col2">
	{{ if .IsSelf }}
		<textarea name="about" rows="6">{{ .About }}</textarea>
	{{ else }}
		{{ .About }}
	{{ end }}
	</div>
</div>
{{ if .IsSelf }}
<div class="row clearfix">
	<div class="col1">Email (private)</div>
	<div class="col2"><input type="email" name="email" value={{ .Email }}></div>
</div>
<div class="row clearfix">
	<div class="col1-offset col2"><input type="submit" value="Update"></div>
</div>
{{ end }}
</form>

<div class="row clearfix">
	<div class="col1-offset col2"><a href="/topics?u={{ .UserName }}">topics</a> {{ if .IsSelf }}(public){{ end }}</div>
</div>
<div class="row clearfix">
	<div class="col1-offset col2"><a href="/comments?u={{ .UserName }}">comments</a> {{ if .IsSelf }}(public){{ end }}</div>
</div>
{{ if .IsSelf }}
<div class="row clearfix">
	<div class="col1-offset col2"><a href="/upvotes?u={{ .UserName }}">upvotes</a> (private)</div>
</div>
{{ end }}

{{ end }}`