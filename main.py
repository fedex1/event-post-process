#!/usr/bin/env python
import sys
import csv

filename = 'move.backup.10Jun2013.EventEntry.csv'
if len(sys.argv) > 1:
    filename = sys.argv[1]
filename2 = 'move.backup.12Jun2013.Event.csv'
if len(sys.argv) > 2:
    filename2 = sys.argv[2]

with open(filename2, 'rb') as csvfile:

    csvreader = csv.reader(csvfile, delimiter=',', quotechar='"')
    #for row in csvreader:
        #print ', '.join(row)

    first = True
    xmap2 = {}
    names = {}
    i = 0
    for row in csvreader:
        if first:
            #print row['address_line_1']
            for item in row:
                xmap2[ item ] = i
                i = i + 1
            #print xmap2 
        else:
            names[row[xmap2['key']]]=row[xmap2['name']]
        first = False
    #print names

#if False:
with open(filename, 'rb') as csvfile:

    csvreader = csv.reader(csvfile, delimiter=',', quotechar='"')
    #for row in csvreader:
        #print ', '.join(row)

    #csvReader = csv.reader(open(filename, 'rb'), delimiter=',', quotechar='"')


    first = True
    map = {}
    map2 = {}
    count1 = {}
    count1withcart = {}
    data = []
    datawithcart = []
    i = 0
    for row in csvreader:
        #print ', '.join(row)
        #print row['cart2']
        #continue
        if first:
            #print row['address_line_1']
            for item in row:
                map[ i ] = item
                map2[ item ] = i
                i = i + 1
            #print map          
            #return

        if True and not first and row[map2['checkout_id']]:
            entry = {}
            entry["created"] = row[map2['created']]
            entry["cart2"] = row[map2['cart2']]
            entry["eventid"] = row[map2['eventid']]
            entry["first"] = row[map2['first']]
            entry["last"] = row[map2['last']]
            entry["age"] = row[map2['age']]
            entry["gender"] = row[map2['gender']]
            entry["checkout_id"] = row[map2['checkout_id']]
            #cart2=str(row[map2['cart2']])
            #print cart2
            data.append(entry)
            if entry["eventid"] not in count1:
                count1[entry["eventid"]] = 1
            else:
                count1[entry["eventid"]] += 1

            entrywithcart = {}
            cart2 = row[map2['cart2']]
            #print cart2
            csvreadercart2 = csv.reader([cart2], delimiter=',', quotechar='"',skipinitialspace=True)
            for rowcart2 in csvreadercart2:
                for item in rowcart2:
                    #print item
                    entrywithcart["cart2new"] = item
                    entrywithcart["created"] = row[map2['created']]
                    entrywithcart["eventid"] = row[map2['eventid']]
                    entrywithcart["first"] = row[map2['first']]
                    entrywithcart["last"] = row[map2['last']]
                    entrywithcart["age"] = row[map2['age']]
                    entrywithcart["gender"] = row[map2['gender']]
                    entrywithcart["checkout_id"] = row[map2['checkout_id']]
                    datawithcart.append(entrywithcart)
                    if item not in count1withcart:
                        count1withcart[item] = 1
                    else:
                        count1withcart[item] += 1

        first = False

data = sorted(data, key=lambda x: x["created"], reverse=True)
data = sorted(data, key=lambda x: x["eventid"], reverse=False)
#print data 
for datum in data:
    #print datum["eventid"], datum["created"]
    #print datum
    pass
#print count1

print '<h2>===Event Entry Overall Count===</h2>'
count2list = []
for key in count1:
    if key not in names:
        #print 'issue: key: >>', key, '<<'
        pass
    else:
        count2 = {}
        count2["name"] = names[key]
        count2["value"] = count1[key]
        #print count1[key], names[key]
        count2list.append(count2)
count2list = sorted(count2list, key=lambda x: x["value"], reverse=True)
# print count2list
print '<table>'
for itemx in count2list:
    print '<tr><td>',itemx["name"],'</td><td>', itemx["value"], '</td></tr>'
print '</table>'


#print datawithcart
#print count1withcart
print '<h2>===Event Entry Detail Count===</h2>'
print '<table>'
for x in count1withcart:
    print '<tr><td>',x,'</td><td>', count1withcart[x], '</td></tr>'
print '</table>'

