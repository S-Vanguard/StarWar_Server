let api = new Vue({
    el: '#api',
    data: {
        username: '',
        input: '',
        select: '',
        selectedTab: 'json',
        jsonSource: 'You\'ve not request any API.',
        isJSONParsed: false,
        loadingJSON: false,
        loadingUser: true,
        parsedJSONTable : [],
        parsingErrorMsg : 'No valid JSON to parse.',
        apiType : 0, // 0 for invalid, 1 for people, 2 for planets, 3 for starships
        tabPosition: 'left',
    },
    methods: {
        handleTabSwitch: function(tab, event) {
            if (tab.name === 'schema' && this.username === '') {
                this.$message.warning('This function is only available for registered users');
                this.parsingErrorMsg = "Login to view the parsed JSON."
            }
        },
        toIndex: function() {
            window.location.href = 'index.html';
        },
        toProfile: function() {
            window.location.href = 'profile.html';
        },
        onSearch: function() {
            if (this.loadingJSON === false) {
                this.loadingJSON = true;
                this.selectedTab = 'json';
                let vueInstance = this;
                axios.get('/' + this.input)
                    .then(function (response) {
                        if (response.status != 200) {
                            vueInstance.$message.error('Server error: ' + response.statusText);
                            vueInstance.jsonSource = 'Error'
                            vueInstance.isJSONParsed = false;
                            vueInstance.parsingErrorMsg = 'No valid JSON to parse.'
                            return;
                        }

                        vueInstance.jsonSource = JSON.stringify(response.data, null, 4).replace(/https:\/\/swapi.co\/api/g, 'http://' + window.location.host)
                        if (vueInstance.username !== '') {
                            vueInstance.isJSONParsed = false;
                            vueInstance.parsingErrorMsg = 'Parsing...'
                            vueInstance.parseJSON();
                            vueInstance.isJSONParsed = true;
                        }
                        else {
                            vueInstance.$message.warning('Login to view the parsed JSON')
                            vueInstance.isJSONParsed = false;
                            vueInstance.parsingErrorMsg = 'Login to view the parsed JSON.';
                        }

                        vueInstance.$message.success('Done')
                    })
                    .catch(function (error) {
                        if (error.response) {
                            vueInstance.$message.error('Connection failed: ' + error.response.statusText);
                            vueInstance.jsonSource = 'Error ' + error.response.status;
                        }
                        else {
                            vueInstance.$message.error('Connection failed: Unknown error');
                            vueInstance.jsonSource = 'Unknown Error';
                        }
                        vueInstance.isJSONParsed = false;
                        vueInstance.parsingErrorMsg = 'No valid JSON to parse.'
                        console.log(error);
                    })
                    .then(function() {
                        vueInstance.loadingJSON = false;
                    });
            }
        },
        logout: function() {
            let vueInstance = this;
            axios.post('/user/logout')
                .then(function(response) {
                    if (response.status != 200) {
                        vueInstance.$message.error('Incorrect status, please try again');
                        return;
                    }

                    if (response.data.status === "OK") {
                        vueInstance.$message.success("Successfully logout")
                        setTimeout(function() {
                            window.location.href = "/";
                        }, 3000)
                    } else if (response.data.status === "Failed") {
                        vueInstance.$message.error(response.data.message);
                        setTimeout(function() {
                            window.location.href = "/";
                        }, 3000)
                    } else {
                        vueInstance.$message.error('Unknown error, please try again');
                    }
                })
                .catch(function (error) {
                    if (error.response) {
                        vueInstance.$message.error('Connection failed: ' + error.response.statusText);
                    }
                    else {
                        vueInstance.$message.error('Connection failed: Unknown error');
                    }
                    console.log(error);
                });
        },
        handleCommand: function(command) {
            switch(command) {
                case 'profile':
                    this.toProfile();
                    break;
                case 'logout':
                    this.logout();
                    break;
            }
        },
        parseJSON: function () {
            let apiTypeKey = this.input.split('/')[0];
            switch(apiTypeKey) {
                case 'people':
                    this.apiType = 1;
                    break;
                case 'planets':
                    this.apiType = 2;
                    break;
                case 'starships':
                    this.apiType = 3;
                    break;
                default:
                    this.apiType = 0;
            }
            if (JSON.parse(this.jsonSource).count !== undefined) {
                // Parse Page of Objects //
                this.parsedJSONTable = JSON.parse(this.jsonSource).results;
            }
            else {
                this.parsedJSONTable = [JSON.parse(this.jsonSource)];
            }
        }
    },
    computed: {
        currentHost: function () {
            return "http://" + window.location.host + '/';
        }
     },
    mounted: function () {
        // Waiting for account module //
        let vueInstance = this;
        axios.get('/user/get')
            .then(function (response) {
                if (response.data.status === 'OK' && response.data.username !== undefined) {
                    vueInstance.$message.success('Welcome, ' + response.data.username);
                    vueInstance.username = response.data.username;
                    vueInstance.loadingUser = false;
                }
                else if (response.data.status === "Failed" && response.data.message !== undefined) {
                    vueInstance.$message.success('Welcome, visitor');
                    vueInstance.username = '';
                    vueInstance.loadingUser = false;
                }
                else {
                    vueInstance.$message.error('Unknown error, try refreshing this page');
                }
            })
            .catch(function (error) {
                if (error.response) {
                    vueInstance.$message.error('Connection failed: ' + error.response.statusText);
                }
                else {
                    vueInstance.$message.error('Connection failed: Unknown error');
                }
                vueInstance.username = '';
                console.log(error)
            });

        // test module //

        // this.$message.success('Welcome, visitor');
        // this.username = '123';
        // this.loadingUser = false;
    }
});