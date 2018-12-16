const emailReg = re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
let main = new Vue({
    el: '#main',
    data: {
        form: {
            username: '',
            password: '',
            confirmPsd: '',
            email: '',
        },     
        isSignUp: false,
    },
    methods: {
        onSubmit: function() {
            let errorMessage = this.check();
            if (errorMessage !== "") {
                this.$message.error(errorMessage);
                return;
            }
            let vueInstance = this;
            if (this.isSignUp === true) {
                axios.post('/user/signUp', {
                    username: vueInstance.form.username,
                    password: vueInstance.form.password,
                    email: vueInstance.form.email
                })
                .then(function (response) {
                    if (response.data.status === "OK") {
                        vueInstance.$message.success("Signed up successfully");
                        vueInstance.isSignUp = false;
                    } else {
                        if (response.data.message !== undefined) {
                            vueInstance.$message.error(response.data.message);
                        } else {
                            vueInstance.$message.error('Unknown error');
                        }
                    }
                })
                .catch(function (error) {
                    if (error.response) {
                        vueInstance.$message.error('Connection failed: ' + error.response.statusText);
                    }
                    else {
                        vueInstance.$message.error('Connection failed: unknown error');
                    }
                    console.log(error);
                });
            }
            else {
                axios.post('/user/signIn', {
                    username: vueInstance.form.username,
                    password: vueInstance.form.password
                })
                .then(function (response) {
                    if (response.data.status === "OK") {
                        vueInstance.$message.success("Signed in successfully");
                        setTimeout(()=>{
                            window.location.href = "/html/swapi.html";
                        }, 3000)
                    } else {
                        if (response.data.message !== undefined) {
                            vueInstance.$message.error(response.data.message);
                        } else {
                            vueInstance.$message.error('Unknown error');
                        }
                    }
                })
                .catch(function (error) {
                    if (error.response) {
                        vueInstance.$message.error('Connection failed: ' + error.response.statusText);
                    }
                    else {
                        vueInstance.$message.error('Connection failed: unknown error');
                    }
                    console.log(error);
                });
            }
        },
        onExchange: function() {
            this.isSignUp = !this.isSignUp;
        },
        toVisitor: function() {
            window.location.href = '/html/swapi.html';
        },
        check: function() {
            if (this.form.username === "") {
                return "Username cannot stay empty";
            }
            if (this.form.password.length < 8) {
                return "Length of password cannot be less than 8 characters";
            }
            if (this.isSignUp && this.form.password !== this.form.confirmPsd) {
                return "Password mismatch";
            }
            if (this.isSignUp && !emailReg.test(this.form.email)) {
                return "Invalid email format";
            }
            return "";
        }
    },
    computed: {},
});