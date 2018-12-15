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
                            vueInstance.$success('Done')
                        }
                        else {
                            vueInstance.$message.warning('Login to view the parsed JSON')
                            vueInstance.isJSONParsed = false;
                            vueInstance.parsingErrorMsg = 'Login to view the parsed JSON.';
                        }
                    })
                    .catch(function (error) {
                        vueInstance.$message.error('Connection failed: ' + error.response.statusText);
                        vueInstance.jsonSource = 'Error ' + error.response.status;
                        vueInstance.isJSONParsed = false;
                        vueInstance.parsingErrorMsg = 'No valid JSON to parse.'
                        console.log(error);
                    })
                    .then(function() {
                        vueInstance.loadingJSON = false;
                    });
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
            if (this.jsonSource.count != undefined) {
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

        // axios.post('/user/get', {})
        //     .then(function (response) {
        //         if (response.data.status === 'OK' && response.data.username !== undefined) {
        //             this.$message.success('Welcome, ' + response.data.username);
        //             this.username = response.data.username;
        //             loadingUser = false;
        //         }
        //         else if (response.data.status === "Failed" && response.data.message !== undefined) {
        //             this.$message.success('Welcome, visitor');
        //             this.username = '';
        //             loadingUser = false;
        //         }
        //         else {
        //             this.$message.error('Unknown error, try refreshing this page');
        //         }
        //     })
        //     .catch(function (error) {
        //         this.$message.error('Connection failed: server does not response');
        //         console.log(error)
        //     })

        // test module //

        this.$message.success('Welcome, visitor');
        this.username = '';
        this.loadingUser = false;
    }
});