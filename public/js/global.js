Vue.component('my-header', {
    data: function() {
        return {}
    },
    methods: {},
    template: '\
    <el-header>\
        <el-row :gutter="20" type="flex" justify="space-between" align="center">\
            <el-col :span="2">\
                <img class="header" src="../img/logo.png">\
            </el-col>\
        </el-row>\
    </el-header>'
});

Vue.component('my-footer', {
    data: function() {
        return  {
            copyright: 'Copyright © 2018 ServiceComputingVanguard. All rights reserved.',
        }
    },
    methods: {
        toGithub: function() {
            window.location.href = 'https://github.com/ServiceComputingVanguard';
        },
    },
    template: '\
    <el-footer>\
        <span>\
            {{ copyright }}\
            <i class="iconfont icon-github" @click="toGithub"></i>\
        </span>\
    </el-footer>',
});
