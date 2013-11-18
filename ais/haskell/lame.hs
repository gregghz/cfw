
--pickMove x y = x ++ " ::: " ++ (foldr (\x acc -> x ++ " | " ++ acc) [] $ words y)

pickMove x y = "49 41"

main = do
	x <- getLine
	let (y,z) = span (/=' ') x
	print $ pickMove y z