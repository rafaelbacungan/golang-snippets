package codewarsSolutions

/*

	A runner, who runs with base speed s with duration t will cover distances d: d = s * t

	However, this runner can sprint for one unit of time with double speed s * 2

	After sprinting, base speed s will permanently be reduced by 1, and for the next one unit of time, the runner will
	enter the recovery phase and can't sprint again.

	Your task, given the base speed s and time t, is to find the maximum possible distance d.

	Example:
	- Given s = 2 and t = 4,
	We could schedule when the runner should sprint so we could get these possible sequences:

	eq.: RRRR
	=> s + s + s + s
	=> 2 + 2 + 2 + 2 = 8
	Seq.: RRRS
	=> s + s + s + s*2
	=> 2 + 2 + 2 + 2*2 = 10
	Seq.: RRSR
	=> s + s + s*2 + (s-1)
	=> 2 + 2 + 2*2 + (2-1) = 9
	Seq.: RSRR
	=> s + s*2 + (s-1) + (s-1)
	=> 2 + 2*2 + (2-1) + (2-1) = 8
	Seq.: RSRS
	=> s + s*2 + (s-1) + (s-1)*2
	=> 2 + 2*2 + (2-1) + (2-1)*2 = 9
	Seq.: SRRR
	=> s*2 + (s-1) + (s-1) + (s-1)
	=> 2*2 + (2-1) + (2-1) + (2-1) = 7
	Seq.: SRRS
	=> s*2 + (s-1) + (s-1) + (s-1)*2
	=> 2*2 + (2-1) + (2-1) + (2-1)*2 = 8
	Seq.: SRSR
	=> s*2 + (s-1) + (s-1)*2 + (s-1-1)
	=> 2*2 + (2-1) + (2-1)*2 + (2-1-1) = 7

	Where:
	- R: Normal Run / Recovery
	- S: Sprint

*/
