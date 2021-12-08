#! usr/bin/perl -wT

use CGI qw(:standard);
use CGI::Carp qw(fatalsToBrowser);
use Time::Local;
use Date::Calc qw(Delta_Days);
use Date::Manip;
use Date::Format;
use strict;

print header;
print start_html('Show Date');
print "<h1>Show Date</h1>";

# 12/31/1999
print time2str("%m/%d/%Y", time);

# 12/02/1999
print time2str("%m/%d/%y", time);

# 12/31/2000
print time2str("%m/%d/%Y", time + (60 * 60 * 24 * 365));

print end_html;

@timery = localtime(time);