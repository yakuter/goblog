<!DOCTYPE html>
<html>

<head>

    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>Globlog | yakuter</title>

    {{template "../includes/head.tpl"}}
    {{ block "pagehead" . }}{{ end }}

</head>

<body class="">

    <div id="wrapper">
        <nav class="navbar-default navbar-static-side" role="navigation">
            <div class="sidebar-collapse">
                <ul class="nav metismenu" id="side-menu">
                    <li class="nav-header">
                        <div class="dropdown profile-element"> <span>
                        <img alt="image" class="img-circle" src="/static/admin/img/profile_small.jpg" />
                            </span>
                            <a data-toggle="dropdown" class="dropdown-toggle" href="#">
                        <span class="clear"> <span class="block m-t-xs"> <strong class="font-bold">{{ .user.UserName }}</strong>
                            </span> <span class="text-muted text-xs block">Art Director <b class="caret"></b></span> </span> </a>
                            <ul class="dropdown-menu animated fadeInRight m-t-xs">
                                <li><a href="profile.html">Profile</a></li>
                                <li><a href="contacts.html">Contacts</a></li>
                                <li><a href="mailbox.html">Mailbox</a></li>
                                <li class="divider"></li>
                                <li><a href="login.html">Logout</a></li>
                            </ul>
                        </div>
                        <div class="logo-element">
                            IN+
                        </div>
                    </li>
                    {{template "../includes/sidebar.tpl" .user}}
                </ul>
            </div>
        </nav>

        <div id="page-wrapper" class="gray-bg">

            {{template "../includes/navbar.tpl"}}

            {{ block "breadcrumb" . }}{{ end }}

            <div class="wrapper wrapper-content animated fadeInRight">
                {{ block "content" . }}{{ end }}               
            </div>
            
            {{template "../includes/footer.tpl"}}
            {{ block "pagefoot" . }}{{ end }}

        </div>
    </div>
    
</body>

</html>