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
    mounted: function() {
        let vueInstance = this;
        axios.get('/user/get')
        .then(function(res) {
            if (res.data.status === 'OK') {
                vueInstance.form.username = res.data.username;
                vueInstance.form.email = res.data.email;
            } else {
                if (res.data.message != undefined) {
                    vueInstance.$message.error(res.data.message);
                } else {
                    vueInstance.$message.error('Unknown error');
                }
            }
        })
        .catch(function (err) {
            vueInstance.$message.error('Connection failed: server does not response');
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
            let vueInstance = this;
            axios.post('/user/update', {
                username: vueInstance.form.username,
                oldPassword: vueInstance.form.oldPsd,
                password: vueInstance.form.password,
                email: vueInstance.isMailChanged ? vueInstance.form.email : ''
            })
            .then(function (res) {
                if (res.data.status === 'OK') {
                    vueInstance.$message.success('Update profile successfully');
                } else {
                    if (res.data.message !== undefined) {
                        vueInstance.$message.error(res.data.message);
                    } else {
                        vueInstance.$message.error('Unknown error');
                    }
                } 
            })
            .catch(function (err) {
                vueInstance.$message.error('Connection failed: server does not response');
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