// eslint-disable-next-line no-unused-vars
function ifnull(txt){
    if(txt.length==0){
        return true;
    }
    var $reg=/\s+/;
    return $reg.test(txt);
}
export function dateFormat () {
    var date = new Date()
    var year = date.getFullYear()
    var month = date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1
    var day = date.getDate() < 10 ? '0' + date.getDate() : date.getDate()
    var hours = date.getHours() < 10 ? '0' + date.getHours() : date.getHours()
    var minutes = date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()
    var seconds = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds()
    let week = date.getDay() // 星期
    let weekArr = [ '星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六' ]

    return year + '年' + month + '月' + day + '日 ' + hours + ':' + minutes + ':' + seconds + ' ' + weekArr[week]
}