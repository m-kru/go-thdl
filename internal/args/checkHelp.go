package args

var checkHelpMsg string = `Check command
=============

Usage
-----

hdl check


Description
-----------

The check command checks for extremely dumb mistakes such as stucking resets
to constant reset value. The check internally consists of independent,
orthogonal scopes. Name of each scope reflects the functional scope that
is actually checked by given scope. Currently following scopes exist:
- clock - checks mistakes related with clock pins mappings,
- process - checks mistakes related with process coding,
- reset - checks mistakes related with reset pins mappings.


Clock scope
-----------

The clock scope is capable of checking following mistakes:

  Mismatched frequency value in port and signal.
    Examples:
      clk_10=>clk_20,
      clock_40_i => clk_160)
      clk70 => clk_80
      clk_70 => clock80_i
      clk70 => clk120
      clk70_i => clk120_i,


Process scope
-------------

The process scope is capable of checking following mistakes:

  Missing sensitivity list in a synchronous process.

    EDA tools correctly synthesize such processes, however one might be surprised
    after including such code into simulation.

  Signal/port used with 'rising_edge()' or 'falling_edge()' function is missing in the sensitivity list

    I have once used 'foo_clk' in the sensitivity list and 'bar_clk' in the 'rising_edge()' function.
    Xilinx Vivado didn't issue even a regular warning. The design was not working correctly,
    and I have lost 3 hours on finding the source of malfunction.


Reset scope
-----------

Note: As hdl is solely based on the text processing and knows nothing about the semantic context,
it imposes some requirements on the port and signal names. Some may find these requirements
stupid and unacceptable. However, they seem to be quite sane if one looks from the lexical point of view.
For example, resets are often associated with some functionality. Let's assume we have reset signal
for resetting some crossbar on a Wishbone bus. To indicate the functionality such signal can be named
{functionality}_{reset} (for example 'wb_rst') or {reset}_{functionality} (for example 'rst_wb').
The hdl reuquires from engineers to use the first form. Why is {reset}_{functionality} wrong?
Because in this case the "reset" part is a verb, such name would be good for procedure or function.
In {functionality}_{reset} the "reset" part is a noun. This order is the valid choice when you realize
that port or signal name is actually a nomina propria. The second requirement is that if 'p' or 'n'
is used to indicate the reset polarity, then it should be placed after the {reset} part.

The reset scope is capable of checking following mistakes:

  Positive reset stuck to '1'.
    Examples:
      rst_p => '1',
      rst_p=>'1',
      reset_p_i=>'1',
      rst => '1',
      reset => '1',
      reset_i => '1',
      reset_p_i => '1',
      RST_P_I=>'1');
      wb_rst_p=> '1',
      wb_rst_p_i=> '1',
      foo_bar_reset => '1',

  Positive reset mapped to negative reset.
    Examples:
      rst_p => rstn,
      rstp => rstn,
      reset => reset_n_i,
      reset_p_i => rst_n);
      wb_rst_p=>  rst_n,
      wb_rst_p_i=> foo_reset_n,

  Positive reset mapped to negated positive reset.
    Examples:
      rst_p => not rst_p_i,
      reset => not(rst_p),
      rst_i => not wb_resetp,

  Negative reset stuck to '0'.
    Examples:
      rst_n => '0',
      rst_n=>'0',
      rstn => '0'
      reset_n_i=>'0',
      reset_n => '0',
      reset_n_i => '0',
      RST_N_I=>'0');
      wb_rst_n=> '0',
      wb_rst_n_i=> '0',
      foo_bar_reset_n => '0',
      foo_rstn=>'0',

  Negative reset mapped to positive reset.
    Examples:
      rst_n => rst,
      wb_reset_n => reset,
      rstn => resetp,
      wb_rst_n => reset);
      foo_rst_n_i => rst_i,

  Negative reset mapped to negated negative reset.
    Examples:
      rst_n => not rst_n_i,
      resetn => not(rst_n),
      rst_i_n => not wb_resetn,
      reset_n_i => not rstn);
`
