const AddDays = (date, n) => {
    const d = new Date(date);
    d.setDate(d.getDate() + n);
    return d.toISOString().split('T')[0];
}

console.log(AddDays('1989-10-15',5));