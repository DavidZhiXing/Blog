#! usr/bin/perl -wT

use CGI qw(:standard);
use CGI::Carp qw(warningsToBrowser fatalsToBrowser);
use Image::Magick;
use strict;
use fcntl qw(:flock);

print header;
print start_html('ImageMagick');
print h1('ImageMagick');

my $image = param("upfile");
unless ($image) {
    print "No image uploaded";
    exit;
} else {
    print "File name: $image";
    my $ctype = uploadInfo($image)->{'Content-Type'};
    print "Content-Type: $ctype";
    # first set the file name
    my $outfile = "./images/outimage.";
    # now determine the file Type
    # depend on what kind of image it is
    if ($ctype =~ m/image\/jpeg/) {
        $outfile .= "jpg";
    } elsif ($ctype =~ m/image\/gif/) {
        $outfile .= "gif";
    } elsif ($ctype =~ m/image\/png/) {
        $outfile .= "png";
    } else {
        print "Unknown file type";
        exit;
    }
    open(OUTFILE, ">$outfile") or die "Can't open $outfile: $!";
    flock(OUTFILE, LOCK_EX) or die "Can't lock $outfile: $!";
    my $file_len = length($image);
    while (read($image, my $buffer, 1024)) {
        print OUTFILE $buffer;
        $file_len += 1024;
        if ($file_len > 1024000) {
            print "File size too large";
            exit;
        }
        close(OUTFILE);
        print "File size: $file_len/1024", "KB<br>\n";
        # now figure out the dimensions of the image
        my($width, $height) = imgsize($outfile);
        print "Width: $width<br>\n";
        print "Height: $height<br>\n";
        print qq(<img src="$outfile" alt="$outfile" width="$width" height="$height">);
        # now create a thumbnail
        my $thumb = Image::Magick->new;
        $thumb->Read($outfile);
        $thumb->Resize(width => 100, height => 100);
        $thumb->Write(filename => "./images/thumb.$outfile");
        print "Thumbnail created<br>\n";
        
        print "File uploaded<br>\n";

    }
    print end_html;

    sub dienice {
        my ($msg) = @_;
        print "Error: $msg<br>\n";
        exit;
    }   
}