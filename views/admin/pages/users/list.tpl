{{ template "../../layouts/master.tpl" . }}

{{ define "breadcrumb" }}
<div class="row wrapper border-bottom white-bg page-heading">
    <div class="col-sm-12">
        <h2>Users</h2>
        <ol class="breadcrumb">
            <li>
                <a href="/admin">Homepage</a>
            </li>
            <li class="active">
                <strong>Users</strong>
            </li>
        </ol>
    </div>
</div>
{{ end }}

{{ define "content" }}
<!--<div class="row">
    <div class="col-lg-12">
        <div class="ibox float-e-margins">
            <div class="ibox-title">
                <h5>Users</h5>
                <div class="ibox-tools">
                    <a href="/admin/users/create" class="btn btn-primary btn-xs">Add New</a>
                </div>
            </div>
            <div class="ibox-content">

                <table class="table">
                    <thead>
                    <tr>
                        <th>#</th>
                        <th>Name Surname</th>
                        <th>Username</th>
                        <th>Email</th>
                        <th>Password</th>
                        <th>Url</th>
                        <th>Info</th>
                        <th>Time</th>
                        <th class="text-right">Action</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $user:=.users}}
                    <tr>
                        <td>{{$user.Id}}</td>
                        <td>{{$user.NameSurname}}</td>
                        <td>{{$user.UserName}}</td>
                        <td>{{$user.Email}}</td>
                        <td>{{$user.Password}}</td>
                        <td>{{$user.Url}}</td>
                        <td>{{$user.Info}}</td>
                        <td>{{$user.RegisterTime}}</td>
                        <td class="text-right">
                            <div class="btn-group">
                                <a href="/admin/users/show/{{$user.Id}}" class="btn btn-xs btn-primary">View</a>
                                <a href="/admin/users/edit/{{$user.Id}}" class="btn btn-xs btn-warning">Edit</a>
                                <a href="/admin/users/delete/{{$user.Id}}" class="btn btn-xs btn-danger">Delete</a>
                            </div>
                        </td>
                    </tr>
                    {{end}}
                    </tbody>
                </table>

            </div>
        </div>
    </div>
</div>-->


<div class="row">
    <div class="col-sm-8">
        <div class="ibox">
            <div class="ibox-title">
                <h5>Users</h5>
                <div class="ibox-tools">
                    <a href="/admin/users/create" class="btn btn-primary btn-xs">Add New</a>
                </div>
            </div>
            <div class="ibox-content">
                <div class="clients-list">
                    {{if .flash.notice}}
                    <div class="alert alert-success alert-dismissable">
                        <button aria-hidden="true" data-dismiss="alert" class="close" type="button">×</button>
                        {{.flash.notice}}.
                    </div>
                    {{ end }}
                    <div class="tab-content">
                        <div id="tab-1" class="tab-pane active">
                            <div class="full-height-scroll">
                                <div class="table-responsive">
                                    <table class="table table-striped table-hover">
                                        <tbody>
                                        {{ if .users }}
                                            {{range $index, $user:=.users}}
                                            <tr>
                                                <td>#{{$user.Id}}</td>
                                                <td class="client-avatar"><img alt="image" src="/static/admin/img/anonymous.png"> </td>
                                                <td><a data-toggle="tab" href="#contact-{{$user.Id}}" class="client-link">{{ $user.NameSurname }}</a></td>
                                                <td>{{$user.UserName}}</td>
                                                <td class="contact-type"><i class="fa fa-envelope"> </i></td>
                                                <td>{{ $user.Email }}</td>
                                                <td class="text-right">

                                                    <a data-toggle="tab" href="#contact-{{$user.Id}}" class="btn btn-xs btn-primary"><i class="fa fa-search"></i></a>
                                                    <a href="/admin/users/edit/{{$user.Id}}" class="btn btn-xs btn-warning"><i class="fa fa-edit"></i></a>
                                                    <a href="/admin/users/delete/{{$user.Id}}" class="btn btn-xs btn-danger"><i class="fa fa-times"></i></a>

                                                </td>
                                            </tr>
                                            {{end}}
                                        {{else}}
                                            <tr>
                                                <td colspan="7">No users here</td>
                                            </tr>
                                        {{end}}
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="col-sm-4">
        <div class="ibox ">
            <div class="ibox-content">
                <div class="tab-content ">
                    {{range $user:=.users}}
                    <div id="contact-{{$user.Id}}" class="tab-pane">
                        <div class="row m-b-lg ">
                            <div class="col-lg-12 text-center">
                                <h2>{{ $user.NameSurname }}</h2>
                                <div class="m-b-sm">
                                    <img alt="image" class="img-circle circle-border m-b-md" src="/static/admin/img/anonymous.png" style="width: 152px">
                                </div>
                            </div>
                            <div class="col-lg-12">
                                <strong>Personal Info</strong><br>
                                <ul class="list-group clear-list">
                                    <li class="list-group-item"><i class="fa fa-vcard-o"> </i>  {{$user.UserName}} <i><small>(#{{$user.Id}})</small></i></li>
                                    <li class="list-group-item"><i class="fa fa-envelope"> </i>  {{$user.Email}}</li>
                                    <li class="list-group-item"><i class="fa fa-link"> </i>  {{$user.Url}}</li>
                                    <li class="list-group-item">{{$user.Info}}</li>
                                </ul>
                            </div>
                        </div>
                    </div>
                    {{end}}
                </div>
            </div>
        </div>
    </div>
</div>

{{ end }}