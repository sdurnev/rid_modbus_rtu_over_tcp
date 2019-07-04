var fs = require('fs');


fs.readFile('1.txt', 'utf8', function (err, data) {
    if (err) throw err;
    var arr = data.split('\n');
    var ansver = [];
    var ansver2 = [];
    console.log(arr.length);
    for (let i = 0; i < arr.length; i++) {
        var tstring = [];
        var tstring2 = [];
        var a;
        var b;
        var c;
        for (let n = 0; n < 6; n++) {
            if (n === 0) {
                tstring.push(parseInt(arr[i + n]));
            } else if (n === 4) {
                tstring.push(parseFloat(arr[i + n]));
                a=parseFloat(arr[i + n]);
                b= arr[i + n];
            } else if (n === 1){
                tstring.push(arr[i + n]);
                c = arr[i + n];
            }else {
                tstring.push(arr[i + n]);
            }
        }
        tstring2.push(a);
        tstring2.push(b+'_'+c);
        ansver2.push(tstring2);
        ansver.push(tstring);
        i = i + 5;
    }
    console.log(JSON.stringify(ansver));
    console.log(JSON.stringify(ansver2));
});
