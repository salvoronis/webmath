var f, curve; // global objects

board = JXG.JSXGraph.initBoard('jxgbox', {
    boundingbox: [-5, 5, 5, -5], axis:true
});
var p = board.create('point',[-3,1],{fixed:true});
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
      plot("x*x-log(x)-2*cos(x)-1");
      break;
    case "2":
      plot("x*x*x-2*x-11");
      break;
    case "3":
      plot("x*x-15")
      break;
    case "4":
      plot("1/(x^(0.5))")
      break;
    case "5":
      plot("-x+3");
      plot("sqrt(-x*x+9)");
      break;
    case "6":
      plot("x*x+1");
      plot("x+1");
      break;
  }
  $.ajax({
    url: '/',
    dataType: 'text',
    data: {error: error, f: f, left:left, right:right, type:type},
    type: 'post',
    success: function(res){
      console.log(res);
      $(".answer").text(res);
    }
  });
});
