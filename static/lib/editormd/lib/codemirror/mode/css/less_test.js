// CodeMirror, copyright (c) by Marijn Haverbeke and others
// Distributed under an MIT license: http://codemirror.net/LICENSE

(function() {
  "use strict";

  var mode = CodeMirror.getMode({indentUnit: 2}, "text/x-less");
  function MT(name) { test.mode(name, mode, Array.prototype.slice.call(arguments, 1), "less"); }

  MT("variable",
     "[variable-1 @base]: [atom #f04615];",
     "[qualifier .class] {",
     "  [property width]: [variable percentage]([number 0.5]); [comment // returns `50%`]",
     "  [property color]: [variable saturate]([variable-1 @base], [number 5%]);",
     "}");

  MT("amp",
     "[qualifier .child], [qualifier .sibling] {",
     "  [qualifier .parent] [atom &] {",
     "    [property color]: [keyword black];",
     "  }",
     "  [atom &] + [atom &] {",
     "    [property color]: [keyword red];",
     "  }",
     "}");

  MT("mixin",
     "[qualifier .mixin] ([variable dark]; [variable-1 @color]) {",
     "  [property color]: [variable darken]([variable-1 @color], [number 10%]);",
     "}",
     "[qualifier .mixin] ([variable light]; [variable-1 @color]) {",
     "  [property color]: [variable lighten]([variable-1 @color], [number 10%]);",
     "}",
     "[qualifier .mixin] ([variable-1 @_]; [variable-1 @color]) {",
     "  [property display]: [atom block];",
     "}",
     "[variable-1 @switch]: [variable light];",
     "[qualifier .class] {",
     "  [qualifier .mixin]([variable-1 @switch]; [atom #888]);",
     "}");

  MT("nest",
     "[qualifier .one] {",
     "  [def @media] ([property width]: [number 400px]) {",
     "    [property font-size]: [number 1.2em];",
     "    [def @media] [attribute print] [keyword and] [property color] {",
     "      [property color]: [keyword blue];",
     "    }",
     "  }",
     "}");
})();
