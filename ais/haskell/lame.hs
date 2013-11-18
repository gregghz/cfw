-- sample input
-- white BR BH BB BQ BK BB BH BR BP BP BP BP BP BP BP BP 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 WP WP WP WP WP WP WP WP WR WH WB WQ WK WB WH WR
-- sample output
-- 49 41

-- get all different moves as ("# #",score), sort by score, return list of just "# #"
genMoves color board = []

-- this version is for a lame quick response that doesn't get past the 2nd round
--pickMove x y = "49 41\n"

pickMove color board = 
	let moves = genMoves color board in
	if length moves > 0 then head moves else "-1 -1" -- if we got nothing then epic fail

-- this version just allows me to verify its interpreting the input properly
--pickMove x y = 
--	let h:rst = words y in
--	x ++ " ::: " ++ (foldr (\x acc -> x ++ " | " ++ acc) h rst) ++ "\n"

main = do
	x <- getLine
	let (y,z) = span (/=' ') x
	putStr $ pickMove y z