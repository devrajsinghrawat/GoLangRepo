var jsonXlsx = require('icg-json-to-xlsx');
var path = require('path');
var filename = path.join('./files', "output.xlsx");


var result = [{
    id: '1',
    name: 'test',
    status: '123'
  },
  {
    id: '2',
    name: 'david',
    status: '323'
  },
  {
    id: '3',
    name: 'ram',
    status: '2323'
  }
];

var outputFile = jsonXlsx.writeFile(filename, JSON.stringify(result));

console.log(outputFile);