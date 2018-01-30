{{ template "/auth/master.tpl" . }}


{{ define "content" }}
<div class="middle-box text-center loginscreen animated fadeInDown">
    <div>
        <div>
            <h1 class="logo-name">gb</h1>
        </div>
        <h3>Welcome to GoBlog</h3>
        <p>This is an open source project developed with golang and beego framework.</p>
        <p>Fill the form to login</p>
        
        <form class="m-t" role="form" method="POST" action="/login">
            {{ .xsrfdata }}
            <div class="form-group">
                <input type="text" class="form-control" name="user_name" placeholder="Username" autofocus required>
            </div>
            <div class="form-group">
                <input type="password" id="password" name="password" class="form-control" placeholder="Password" required>
            </div>
            <button type="submit" class="btn btn-primary block full-width m-b">Login</button>

            <a href="#"><small>Did you forget your password?</small></a>
            <p class="text-muted text-center"><small>Don't you have an account=</small></p>
            <a class="btn btn-sm btn-white btn-block" href="/register">Create Account</a>
        </form>
        <p class="m-t"> <small>Coded with love by yakuter.</small> </p>
    </div>
</div>
{{ end }}