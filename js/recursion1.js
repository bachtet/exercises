// Towers of Hanoi, print out length.

const towers = [[5,4,3,2,1], [], []];

move = (n, a, b, c) => {
	if(n <= 0) {return;}
  move(n-1, a, c, b);
  towers[c].push(towers[a].pop());
  console.log(`tower a: ${towers[0].length}, tower b: ${towers[1].length}, tower c: ${towers[2].length}`);
  move(n-1, b, a, c)
}

move(5, 0, 1, 2);