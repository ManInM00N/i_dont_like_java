function ifnull(txt){
    if(txt.length==0){
        return true;
    }
    var $reg=/\s+/;
    return $reg.test(txt);
}
window.onload(()=>{
    $("#email").blur(function(){
        if (ifnull($("#email").val())){
            $("#email_msg").html("<p style='color: red; border:none ;margin-top:2px ;margin-bottom: 2px;width: 100%'>邮箱不能为空</p>");
        }else {
            var reg = /\w[-.\w]*\@[-a-z0-9]+(\.[-a-z0-9]+)*\.(com|cn|edu|uk)/ig;
            if (!reg.test($("#email").val())){
                $("#email_msg").html("<p style='color: red ;border:none ;margin-top:2px ;margin-bottom: 2px;width: 100%'>邮箱格式错误</p>");
            }else {
                $("#email_msg").html("<p style='color: greenyellow;border: none;float: right'>√</p>");
            }
        }
    });
    $("#username").blur(function(){
        if(ifnull($("#username").val())){
            $("#uname").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width: 100%'>用户名不能为空</p>");
        }else{
            // var $reg=/^[\w\u4E00-\u9FA5]{4,20}$/;
            var $reg =/([\s\S]*)/;

            if(!$reg.test($("#username").val())){
                $("#uname").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>用户名至少要4-20位</p>");
            }else{
                $("#uname").html("<p style='color: greenyellow;border: none;float: right'>√</p>");

            }
        }
    });
    $("#passwordfirst").blur(function(){
        if(ifnull($("#passwordfirst").val())){
            $("#passwordfir").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>密码不能为空</p>");
        }else{
            var $reg=/^\w{6,}$/;
            if(!$reg.test($("#passwordfirst").val())){
                $("#passwordfir").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>密码至少要6位,且由数字及字母组成</p>");
            }else{
                $("#passwordfir").html("<p style='color: greenyellow;border: none;float: right'>√</p>");
            }
        }
    });
    $("#passwordsecond").blur(function(){
        if(ifnull($(this).val())){
            $("#passwordsec").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>两次密码不一致</p>");
        }else{
            if($("#passwordfirst").val()!=$("#passwordsecond").val()){
                $("#passwordsec").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>两次密码不一致</p>");
            }else{
                $("#passwordsec").html("<p style='color: greenyellow;border: none;float: right'>√</p>");
            }
        }
    });
})

function check(){
    if(ifnull($("#username").val())){
        alert("用户名为空！")
        $("#uname").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>用户名不能为空</p>");
        return false;
    }else{
        // var $reg=/^[\w\u4E00-\u9FA5]{4,20}$/;
        var $reg =/([\s\S]*)/;
        if(!$reg.test($("#username").val())){

            alert("用户名格式不正确！");
            $("#uname").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>用户名至少要4-20位</p>");
            return false;
        }
    }
    if(ifnull($("#passwordfirst").val())){
        alert("用户名为空！")

        $("#passwordfir").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>密码不能为空</p>");
        return false;
    }else{
        var $reg=/^\w{6,}$/;
        if(!$reg.test($("#passwordfirst").val())){
            alert("密码格式不正确！")

            $("#passwordfir").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>密码至少要6位,且由数字及字母组成</p>");
            return false;
        }
    }
    if(ifnull($(this).val())){
        alert("未二次确认密码！")

        $("#passwordsec").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>两次密码不一致</p>");
        return false;
    }else{
        if($("#passwordfirst").val()!=$("#passwordsecond").val()){
            alert("两次密码不一致")

            $("#passwordsec").html("<br><p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>两次密码不一致</p>");
            return false;
        }
    }
    if (ifnull($("#email").val())){
        alert("邮箱为空！")
        $("#email_msg").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>邮箱不能为空</p>");
        return false;
    }else {
        var reg =/\w[-.\w]*\@[-a-z0-9]+(\.[-a-z0-9]+)*\.(com|cn|edu|uk)/ig;
        if (reg.test($(this).val())){
            alert("邮箱格式错误！")
            $("#email_msg").html("<p style='color: red; border:none; margin-top:2px ;margin-bottom: 2px;width:100%'>邮箱格式错误</p>");
            return  false;
        }
    }

    return true;
}