
--pickMove x y = x ++ " ::: " ++ (foldr (\x acc -> x ++ " | " ++ acc) [] $ words y)

-- sample input
-- white BR BH BB BQ BK BB BH BR BP BP BP BP BP BP BP BP 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 WP WP WP WP WP WP WP WP WR WH WB WQ WK WB WH WR

pickMove x y = "49 41\n"

main = do
	x <- getLine
	let (y,z) = span (/=' ') x
	putStr $ pickMove y z