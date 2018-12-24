
jQuery(document).ready(function () {

    /*
        Fullscreen background
    */
    $.backstretch("assets/img/backgrounds/1.jpg");

    /*
        Form validation
    */
    $('.login-form input[type="text"], .login-form input[type="password"], .login-form textarea').on('focus', function () {
        $(this).removeClass('input-error');
    });

    $('.btn-submit').click(function (e) {
        var ok = true;
        $(".login-form").find('input[type="text"], input[type="password"], textarea').each(function () {
            if ($(this).val() == "") {
                ok = false;
                $(this).addClass('input-error');
            }
            else {
                $(this).removeClass('input-error');
            }
        });
        if (ok) {
            e.preventDefault();
        } else {
            var json = JSON.stringify($(".login-form").serializeObject());
            alert(json);
            /*
                $.ajax({
                    url: "",
                    type: "POST",
                    data: jsonuserinfo,
                    contentType: "application/json",  //缺失会出现URL编码，无法转成json对象
                    success: function () {
                        alert("send");
                    }
                });
            */
        }
    });

});


$.fn.serializeObject = function () {
    var o = {};
    var a = this.serializeArray();
    $.each(a, function () {
        if (o[this.name]) {
            if (!o[this.name].push) {
                o[this.name] = [o[this.name]];
            }
            o[this.name].push(this.value || '');
        } else {
            o[this.name] = this.value || '';
        }
    });
    return o;
};