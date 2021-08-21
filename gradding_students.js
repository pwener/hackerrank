// grade - round(grade) % 5 < 3, round(grade)
// grade < 38, 38

const findNextMultipleOfFive = (n) => {
    if (n % 5 === 0) {
        return n;
    }
    return findNextMultipleOfFive(n + 1);
}

function gradingStudents(grades) {    
    return grades.map((g) => {
        if (g < 38) {
            return g;
        }
        const nextM = findNextMultipleOfFive(g);        

        return nextM - g < 3 ? nextM : g;
    })
}
