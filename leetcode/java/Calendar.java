public class Calendar
{
    private Calendar(){}

    public Calendar getInstance()
    {
        return new Calendar();
    }

    public void set(int year, int month, int day)
    {
    }

    public void setTime(Date date)
    {
        _date = date;
    }

    public void add(int field, int amount)
    {
        switch (field)
        {
            case Calendar.YEAR:
                _date.setYear(_date.getYear() + amount);
                break;
            case Calendar.MONTH:
                _date.setMonth(_date.getMonth() + amount);
                break;
            case Calendar.DATE:
                _date.setDate(_date.getDate() + amount);
                break;
            case Calendar.HOUR:
                _date.setHours(_date.getHours() + amount);
                break;
            case Calendar.MINUTE:
                _date.setMinutes(_date.getMinutes() + amount);
                break;
            case Calendar.SECOND:
                _date.setSeconds(_date.getSeconds() + amount);
                break;
            case Calendar.MILLISECOND:
                _date.setMilliseconds(_date.getMilliseconds() + amount);
                break;
            default:
                throw new IllegalArgumentException("Invalid field: " + field);
        }
    }

    public Date getTime()
    {
        return _date;
    }

    private Date _date;

}