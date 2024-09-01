defmodule Seqn do
  def diffs(seq) do
    seq
    |> Enum.chunk_every(2, 1, :discard)
    # |> IO.inspect(label: "diffs")
    |> Enum.map(fn [a, b] -> b - a end)
  end

  def get_next_in_seq(seq) do
    cond do
      Enum.all?(seq, fn x -> x == 0 end) -> 0
      true -> Enum.at(seq, -1) + (seq |> diffs() |> get_next_in_seq())
    end
  end
end

seqs =
  IO.stream(:stdio, :line)
  |> Enum.map(
	fn line ->
		line
		|> String.split()
		|> Enum.map(
			fn numStr ->
				numStr
				|> Integer.parse()
				|> elem(0)
    		end
		)
  	end
  )

# IO.inspect(seqs, label: "Sequences")

IO.write("\nPart 1\n")

seqs
|> Enum.map(fn seq -> Seqn.get_next_in_seq(seq) end)
# |> Enum.map(fn s -> IO.inspect(s) end)
|> Enum.sum()
|> IO.inspect(label: "Sum of next vals")


IO.write("\nPart 2\n")

seqs
|> Enum.map(&Enum.reverse/1)
|> Enum.map(&Seqn.get_next_in_seq/1)
# |> Enum.map(&IO.inspect/1)
|> Enum.sum()
|> IO.inspect(label: "Sum of prev vals")
