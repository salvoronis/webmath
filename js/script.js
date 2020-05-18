var f, curve; // global objects

board = JXG.JSXGraph.initBoard('jxgbox', {
    boundingbox: [-5, 5, 5, -5], axis:true
});
//var p = board.create('point',[-3,1],{fixed:true});
function plot(graph){
  f = board.jc.snippet(graph, true, 'x', true);
  curve = board.create('functiongraph',[f,
    function(){
      var c = new JXG.Coords(JXG.COORDS_BY_SCREEN,[0,0],board);
      return c.usrCoords[1];
    },
    function(){
      var c = new JXG.Coords(JXG.COORDS_BY_SCREEN,[board.canvasWidth,0],board);
      return c.usrCoords[1];
    }
  ]);
}

function clearAll() {
  JXG.JSXGraph.freeBoard(board);
  board = JXG.JSXGraph.initBoard('jxgbox', {boundingbox:[-5, 5, 5, -5], axis:true});
  f = null;
  curve = null;
}

function doDot(x,y){
  var p = board.create('point',[parseFloat(x),parseFloat(y)],{fixed:true});
}



function openCity(evt, cityName) {
    var i, tabcontent, tablinks;
    tabcontent = document.getElementsByClassName("tabcontent");
    for (i = 0; i < tabcontent.length; i++) {
        tabcontent[i].style.display = "none";
    }
    tablinks = document.getElementsByClassName("tablinks");
    for (i = 0; i < tablinks.length; i++) {
        tablinks[i].className = tablinks[i].className.replace(" active", "");
    }
    document.getElementById(cityName).style.display = "block";
    evt.currentTarget.className += " active";
}

$(".tabcontent").on("submit", function(event){
  event.preventDefault();
  clearAll();
  var error = $(this).find("input[name='error']").val();
  var f = $(this).find("select[name='function']").find(":selected").val();
  var left = $(this).find("input[name='left']").val();
  var right = $(this).find("input[name='right']").val();
  var type = $(this).find("input[name='type']").val();
  switch (f) {
    case "1":
      plot("log(x)");
      break;
    case "2":
      plot("x*x*x-2*x-11");
      break;
    case "3":
      plot("x*x-15")
      break;
    case "4":
      plot("1/(x^(0.5))-0.5")
      break;
    case "5":
      plot("sin(x)");
      plot("(x-2)^(1/2)");
      plot("-(x-2)^(1/2)");
      break;
    case "6":
      plot("sin(x)");
      plot("(4-x)^(1/3)");
      plot("-(-4+x)^(1/3)");
      break;
  }
  $.ajax({
    url: '/',
    dataType: 'text',
    data: {error: error, f: f, left:left, right:right, type:type},
    type: 'post',
    success: function(res){
      console.log(res);
      var answers = res.split(";");
      doDot(answers[0],answers[1]);
      $(".answer").text("x: "+answers[0]+" y: "+answers[1]);
    }
  });
});
