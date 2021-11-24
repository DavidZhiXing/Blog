typedef struct LVisionParam{
    double* intrisic;
    double* distortion;
    double* rotation;
    double* translation;
    double* projection;

    unsigned int marker_size;
    void* marker_data;

    int image_width;
    int image_height;
    bool is_enable_low_cost;
    bool is_enable_debug;

    bool bDraw;
    bool bDraw_marker;
    double fov_width;
    double fov_height;
} LVisionParam;

/* C# TO C++ CONVERTER NOTE: The following .NET attribute has no direct equivalent in native C++:
'[StructLayout(LayoutKind.Sequential)]' converted to struct 

 [StructLayout(LayoutKind.Sequential)]
 public struct LVisionParam
 {
     public double[] intrisic;
     public double[] distortion;
     public double[] rotation;
     public double[] translation;
     public double[] projection;

     public int marker_size;
     public IntPtr marker_data;

     public int image_width;
     public int image_height;
     public bool is_enable_low_cost;
     public bool is_enable_debug;

     public bool bDraw;
     public bool bDraw_marker;
     public double fov_width;
     public double fov_height;
 }

 or 

[StructLayout(LayoutKind.explicit)]
public struct LVisionParam
{
    [FieldOffset(0)]
    public double[] intrisic;
    [FieldOffset(4)]
    public double[] distortion;
    [FieldOffset(8)]
    public double[] rotation;
    [FieldOffset(12)]
    public double[] translation;
    [FieldOffset(16)]
    public double[] projection;

    [FieldOffset(20)]
    public int marker_size;
    [FieldOffset(24)]
    public IntPtr marker_data;

    [FieldOffset(28)]
    public int image_width;
    [FieldOffset(32)]
    public int image_height;
    [FieldOffset(36)]
    public bool is_enable_low_cost;
    [FieldOffset(37)]
    public bool is_enable_debug;

    [FieldOffset(38)]
    public bool bDraw;
    [FieldOffset(39)]
    public bool bDraw_marker;
    [FieldOffset(40)]
    public double fov_width;
    [FieldOffset(44)]
    public double fov_height;
}

*/

