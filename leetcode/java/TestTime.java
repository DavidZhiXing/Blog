public class Test{

    public void AddHoursToDate(Date date, int hours){
        Calendar cal = Calendar.getInstance();
        cal.setTime(date);
        cal.add(Calendar.HOUR, hours);
        date = cal.getTime();
    }

    public void AddHoursToDateTime(DateTime dateTime, int hours){
        dateTime = dateTime.plusHours(hours);
    }

    public void AddHourToTime(Time time, int hours){
        time = time.plusHours(hours);
    }
}