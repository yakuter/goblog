{{ template "/auth/master.tpl" . }}

{{ define "content" }}
<div class="middle-box text-center loginscreen animated fadeInDown">
    <div>
        <div>
            <h1 class="logo-name">gb</h1>
        </div>
        <h3>Register!</h3>
        <p>Fill the form to register.</p>
        <form class="m-t" role="form" method="POST" action="/register">
            {{ .xsrfdata }}

            <div class="form-group">
                <input type="text" name="name_surname" id="name_surname" class="form-control" placeholder="Name Surname" autofocus>
            </div>

            <div class="form-group">
                <input type="text" name="user_name" id="user_name" class="form-control" placeholder="Username *" required>
            </div>
            
            <div class="form-group">
                <input type="email" name="email" id="email" class="form-control" placeholder="E-mail *" required>
            </div>
            
            <div class="form-group">
                <input type="password" name="password" id="password" class="form-control" placeholder="Password *"  required>
            </div>
            
            <div class="form-group">
                <input type="password" name="password_confirm" id="password-confirm" class="form-control" placeholder="Password Confirm *"  required>
            </div>

            <button type="submit" class="btn btn-primary block full-width m-b">Register</button>

            <p class="text-muted text-center"><small>Do you already have an account?</small></p>
            <a class="btn btn-sm btn-white btn-block" href="/login">Login</a>
        </form>
        <p class="m-t"> <small>Coded with love by yakuter.</small> </p>
    </div>
</div>
{{ end }}