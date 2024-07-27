//
// Pocket version of Axelrod's Iterated Prisoner's Dilemma Tournament
//
//		C	D
//	C	(1,1)	(3,0)
//	D	(0,3)	(2,2)
//
//	A:C	B:C	(1,1)
//	A:C	B:D	(0,3)
//	A:D	B:C	(3,0)
//	A:D	B:D	(2,2)
//

var PAYOFFS = {
  'CC': [3, 3],
  'CD': [0, 5],
  'DC': [5, 0],
  'DD': [1, 1]
}

// var PAYOFFS = { 
// 	'CC': [1,1],
// 	'CD': [3,0],
// 	'DC': [0,3],
// 	'DD': [2,2]
// }

// var PAYOFFS = { 
// 	'CC': [2,2],
// 	'CD': [-1,4],
// 	'DC': [4,-1],
// 	'DD': [1,1]
// }

function always_cooperate() {
  this.name = 'Always Cooperate';
  this.last_op_move = '';
  this.move = function() { return ('C') };
}

function always_defect() {
  this.name = 'Always Defect';
  this.last_op_move = '';
  this.move = function() { return ('D') };
}

function random() {
  this.name = 'Random 50/50';
  this.last_op_move = '';
  this.move = function() {
    if (Math.random() > 0.5) {
      return ('C');
    } else {
      return ('D');
    }
  };
}

function mostly_cooperate() {
  this.name = 'Cooperate in 75%';
  this.last_op_move = '';
  this.move = function() {
    if (Math.random() > 0.25) {
      return ('C');
    } else {
      return ('D');
    }
  };
}

function mostly_defect() {
  this.name = 'Defect in 75%';
  this.last_op_move = '';
  this.move = function() {
    if (Math.random() > 0.25) {
      return ('D');
    } else {
      return ('C');
    }
  };
}

// Try to defect after the opponent cooperates 3 times in a row.
// If the opponent is a simpleton, defect further. If not, switch to tit-for-tat
function opportunist() {
  this.name = 'Opportunist';
  this.last_op_move = '';
  this.coops_so_far = 0;
  this.coop_threshold = 3;
  this.move = function() {
    if (this.last_op_move == 'C') {
      this.coops_so_far++;
    } else {
      this.coops_so_far = 0;
    }

    if (this.coops_so_far >= this.coop_threshold) {
      return ('D');
    } else {
      if (this.last_op_move == 'D') {
        return ('D');
      } else {
        return ('C');
      }
    }
  };
}

// Do what your opponent did last time
function t4t() {
  this.name = 'Tit for Tat';
  this.last_op_move = '';
  this.move = function() {
    if (this.last_op_move == '' || this.last_op_move == 'C') {
      return ('C');
    } else {
      return ('D');
    }
  };
}

// Do as your opponent did but stop cooperating after his first defection
function unforgiving_t4t() {
  this.name = 'Unforgiving Tit for Tat';
  this.last_op_move = '';
  this.flag = true;
  this.move = function() {
    if (this.flag) {
      if (this.last_op_move == 'D') {
        this.flag = false;
        return ('D');
      }
      return ('C');
    } else {
      return ('D');
    }
  };
}

// Defect back only on second transgression in a row
function t42t() {
  this.name = 'Tit for Two Tats';
  this.last_op_move = '';
  this.flag = false;
  this.move = function() {
    if (this.last_op_move == '' || this.last_op_move == 'C') {
      this.flag = false;
      return ('C');
    } else {
      if (this.flag) {
        this.flag = false;
        return ('D');
      } else {
        this.flag = true;
        return ('C');
      }
    }
  };
}

// Allow up to two defections throughout the game
function fool_me_twice() {
  this.name = 'Fool me twice';
  this.last_op_move = '';
  this.counter = 0;
  this.tolerance = 2;
  this.move = function() {
    if (this.last_op_move == '' || this.last_op_move == 'C') {
      if (this.counter >= this.tolerance) {
        return ('D');
      } else {
        return ('C');
      }
    } else {
      if (this.counter >= this.tolerance) {
        this.counter++;
        return ('D');
      } else {
        this.counter++;
        return ('C');
      }
    }
  };
}

// Allow up to three defections throughout the game
function fool_me_thrice() {
  var self = new fool_me_twice();
  self.tolerance = 3;
  self.name = 'Fool me thrice';
  return (self);
}

// function fool_me_thrice () {
// 	this.name = 'Fool me thrice';
// 	this.last_op_move = '';
// 	this.counter = 0;
// 	this.tolerance = 3;
// 	this.move = function () {
// 		if (this.last_op_move == '' || this.last_op_move == 'C') {
// 			if (this.counter >= this.tolerance) {
// 				return ('D');
// 			} else {
// 				return ('C');
// 			}
// 		} else {
// 			if (this.counter >= this.tolerance) {
// 				this.counter++;
// 				return ('D');
// 			} else {
// 				this.counter++;
// 				return ('C');
// 			}
// 		}
// 	};
// }

// Establish the probability of cooperation and defection from past moves
// in the game and make the move mirroring said probability
function e4e() {
  this.name = 'Eye for an Eye';
  this.last_op_move = '';
  this.d_counter = 0;
  this.counter = 0;
  this.move = function() {
    if (this.last_op_move == '' || this.last_op_move == 'C') {
      if (this.d_counter == 0) {
        this.counter++;
        return ('C');
      } else {
        var prob = this.d_counter / this.counter;
        if (Math.random() > prob) {
          this.counter++;
          return ('C');
        } else {
          this.counter++;
          return ('D');
        }
      }
    } else {
      this.d_counter++;
      var prob = this.d_counter / this.counter;
      if (Math.random() > prob) {
        this.counter++;
        return ('C');
      } else {
        this.counter++;
        return ('D');
      }
    }
  };
}

// ========================================= Let's roll =========================================

var results = {};
player_list = [always_cooperate, always_defect, mostly_cooperate, mostly_defect, t4t, unforgiving_t4t, t42t, random, opportunist, fool_me_twice, fool_me_thrice, e4e];

for (i = 0; i < player_list.length; i++) {
  for (j = i; j < player_list.length; j++) {

    var player1 = new player_list[i];
    if (!results[player1.name])
      results[player1.name] = 0;
    var player2 = new player_list[j];
    if (!results[player2.name])
      results[player2.name] = 0;

    print("Pairing", player1.name, "vs", player2.name);

    var p1tmp = 0;
    var p2tmp = 0;
    var n = 100;

    while (n > 0) {
      p1move = player1.move();
      p2move = player2.move();
      outcome = (PAYOFFS[p1move + p2move]);
      player1.last_op_move = p2move;
      player2.last_op_move = p1move;
      p1tmp += outcome[0];
      p2tmp += outcome[1];
      // 			print (p1move, p2move, outcome);
      n--;
    }

    print(p1tmp, p2tmp);

    if (player1.name == player2.name) {
      results[player1.name] += (p1tmp + p2tmp) / 2;
    } else {
      results[player1.name] += p1tmp;
      results[player2.name] += p2tmp;
    }

  }
}

// ========================================= Results =========================================

print('------------------------------');

for (i in results) {
  print(i, ":", results[i]);
}
