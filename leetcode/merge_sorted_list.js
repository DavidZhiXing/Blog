const mergeTwoLists = function(l1, l2) {
    if(l1 === null) return l2;
    if(l2 === null) return l1;
    if(l1.val < l2.val){
        l1.next = mergeTwoLists(l1.next, l2);
        return l1;
    }else{
        l2.next = mergeTwoLists(l1, l2.next);
        return l2;
    }
};



// test
const l1 = {val: 1, next: {val: 2, next: {val: 4, next: null}}};
const l2 = {val: 1, next: {val: 3, next: {val: 4, next: null}}};
console.log(mergeTwoLists(l1, l2));