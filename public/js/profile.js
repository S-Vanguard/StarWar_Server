const emailReg = re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
let profile = new Vue({
    el: '#profile',
    data: {
        labelPosition: 'left',
        form: {
            username: 'user',
            oldPsd: '',
            password: '',
            confirmPsd: '',
            email: '',
        },
        isMailChanged: false,
    },
    created: function() {
        axios.get('/user/get')
        .then(function(res) {
            if (res.data.status === 'OK') {
                this.form.username = res.data.username;
                this.form.email = res.data.email;
            } else {
                if (res.data.message != undefined) {
                    this.$message.error(res.data.message);
                } else {
                    this.$message.error('Unknown error');
                }
            }
        })
        .catch(function (err) {
            this.$message.error('Connection failed: server does not response');
            console.log(err)
        });
    },
    methods: {
        mailChange: function(val) {
            this.isMailChanged = true;
        },
        onSubmit: function() {
            let errMsg = this.check();
            if (errMsg !== '') {
                this.$message.error(errMsg);
                return;
            }
            axios.post('/user/update', {
                username: this.form.username,
                oldPassword: this.form.oldPsd,
                password: this.form.password,
                email: this.isMailChanged ? this.form.email : ''
            })
            .then(function (res) {
                if (res.data.status === 'OK') {
                    this.$message.success('Update profile successfully');
                } else {
                    if (res.data.message !== undefined) {
                        this.$message.error(res.data.message);
                    } else {
                        this.$message.error('Unknown error');
                    }
                } 
            })
            .catch(function (err) {
                this.$message.error('Connection failed: server does not response');
                console.log(err)
            });
        },
        check: function() {
            if (this.form.password !== this.form.confirmPsd) {
                return 'Password mismatch';
            }
            if (this.isMailChanged && !emailReg.test(this.form.email)) {
                return 'Invalid email format';
            }
            return '';
        }
    },
    computed: '',
});