{{ template "../layouts/master.tpl" . }}

{{ define "user" }}
{{ .user.UserName }}
{{ end }}

{{ define "breadcrumb" }}
<div class="row wrapper border-bottom white-bg page-heading">
    <div class="col-sm-4">
        <h2>Dashboard</h2>
        <ol class="breadcrumb">
            <li class="active">
                <strong>Homepage</strong>
            </li>
        </ol>
    </div>
    <div class="col-sm-8">
        <div class="title-action">
            {{str2html .word}}
        </div>
    </div>
</div>
{{ end }}

{{ define "content" }}

<div class="row">
    <div class="col-lg-12">

        <div class="ibox">

                <div class="col-lg-5">
                    <h2>About Goblog</h2>
                    <p><a href="#">Yakuter's Goblog</a> is an open source project written with <a href="https://golang.org/">Golang</a> and <a href="https://beego.me/">Beego Framework</a> used in the background.</p>
                </div>

                <div class="col-lg-7">
                    <h2>What is Next</h2>
                    <table class="table table-borderless">
                        <tbody>
                        <tr>
                            <td><a href="/admin/categories/create"><i class="fa fa-folder-open"></i> Add a new category</a></td>
                            <td><a href="/admin/posts/create"><i class="fa fa-folder-open"></i> Write a new awesome blog post</a></td>
                        </tr>
                        <tr>
                            <td><a href="/admin/pages/create"><i class="fa fa-folder-open"></i> Add a new page</a></td>
                            <td><a href="/admin/users"><i class="fa fa-folder-open"></i> Manage users</a></td>
                        </tr>
                        </tbody>
                    </table>
                </div>


        </div>
    </div>
</div>

<div class="row">
    <div class="col-lg-5">
        <div class="ibox float-e-margins">
            <div class="ibox-title">
                <h5>At a Glance</h5>
            </div>
            <div class="ibox-content">

                <table class="table table-borderless">
                    <tbody>
                        <tr>
                            <td><a href="/admin/categories"><i class="fa fa-folder-open"></i> {{.catNumber}} Categories</a></td>
                            <td><a href="/admin/posts"><i class="fa fa-folder-open"></i> {{.postNumber}} Posts</a></td>
                        </tr>
                        <tr>
                            <td><a href="/admin/pages"><i class="fa fa-folder-open"></i> {{.pageNumber}} Pages</a></td>
                            <td><a href="/admin/users"><i class="fa fa-folder-open"></i> {{.userNumber}} Users</a></td>
                        </tr>
                    </tbody>
                </table>

            </div>
        </div>
    </div>
    <div class="col-lg-7">
        <div class="ibox float-e-margins">
            <div class="ibox-title">
                <h5>Recent Posts</h5>
            </div>
            <div class="ibox-content">

                <table class="table table-borderless">
                    <tbody>
                    {{ if .recentPosts}}
                        {{ range $recent:=.recentPosts}}
                            <tr>
                                <td>{{date $recent.Time "d.m.Y / H:i:s"}}</td>
                                <td><a href="/admin/posts/edit/{{$recent.Id}}">{{$recent.Title}}</a></td>
                            </tr>
                        {{end}}
                    {{end}}
                    </tbody>
                </table>

            </div>
        </div>
    </div>
</div>


{{ end }}