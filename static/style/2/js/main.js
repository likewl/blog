new Vue({
    el: "#demo",
    data: {
        result1: null,
        result2: null,
        input2: ""
    },
    created: function() {
        // the created hook is called after the instance is created
    },
    methods: {
        test1: function () {
            this.result1 = this.$refs.input1.value + " 成功!";
        },
        test2: function () {
            this.result2 = this.input2 + " 成功!";
        }
    }
});


$("#search_btn").click(function() {
  var search_input = $("#search_input").val();
  var baidu_url = '/research?research=' + encodeURIComponent(search_input);
  window.open(baidu_url);
});

$("#search_input").keyup(function(event) {
    if (event.keyCode === 13) {
        $("#search_btn").click();
    }
});

$("#search_btn_phone").click(function() {
    var search_input = $("#search_phone").val();
    var baidu_url = '/research?research=' + encodeURIComponent(search_input);
    window.open(baidu_url);
});
$("#search_phone").keyup(function(event) {
    if (event.keyCode === 13) {
        $("#search_btn").click();
    }
});


var captureOutboundLink = function(url) {
   ga('send', 'event', 'outbound', 'click', url, {
     'transport': 'beacon',
     'hitCallback': function(){document.location = url;}
   });
}
