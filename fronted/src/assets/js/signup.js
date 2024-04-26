// eslint-disable-next-line no-unused-vars
function ifnull(txt){
    if(txt.length==0){
        return true;
    }
    var $reg=/\s+/;
    return $reg.test(txt);
}
// eslint-disable-next-line no-unused-vars
function checksame(rule ,value ,callback){
    if (value === '') {
        callback(new Error('请再次输入密码'))
        // password 是表单上绑定的字段
    } else if (value !== this.info.password) {
        callback(new Error('两次输入密码不一致!'))
    } else {
        callback()
    }
}