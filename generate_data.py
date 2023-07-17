#!/usr/bin/env python3

import re


male_data = {}
female_data = {}
with open("./firstnames.csv", "r") as f:
    lines = f.readlines()
    for line in lines:
        parts = line.split(",")
        sex = parts[3].strip('"')
        name = parts[4].strip('"')
        indicator = parts[5].strip('"')
        value = parts[12].strip('"')
        if indicator == 'Proportion (%)':
            if sex == 'Male':
                if name not in male_data:
                    male_data[name] = 0
                male_data[name] += float(value)
            elif sex == 'Female':
                if name not in female_data:
                    female_data[name] = 0
                female_data[name] += float(value)

surnames = {}
# name,rank,count,prop100k,cum_prop100k,pctwhite,pctblack,pctapi,pctaian,pct2prace,pcthispanic
with open("./surnames.csv", "r") as f:
    lines = f.readlines()
    for line in lines[1:]:
        parts = line.split(",")
        name = parts[0].strip('"')
        count = parts[2].strip('"')
        if name not in surnames:
            surnames[name] = 0
        surnames[name] += float(count)

ages = {}
totalAgeWeights = 0
# "REF_DATE","GEO","DGUID","Sex","Age group","UOM","UOM_ID","SCALAR_FACTOR","SCALAR_ID","VECTOR","COORDINATE","VALUE","STATUS","SYMBOL","TERMINATED","DECIMALS"
with open("./population_age.csv", "r") as f:
    lines = f.readlines()
    for line in lines[1:]:
        parts = line.split(",")
        geo = parts[1].strip('"')
        gender = parts[3].strip('"')
        age = parts[4].strip('"')
        value = parts[11].strip('"')
        if value == "":
            continue
        if "and over" in age:
            continue
        if gender != "Both sexes":
            continue
        if geo != "Canada":
            continue
        if not re.match(r'\d+ years', age):
            continue

        age = int(age.split(" ")[0])
        if age in ages:
            continue
        ages[age] = int(value)
        totalAgeWeights += int(value)

with open("./faker/data/ages.go", "w") as f:
    f.write("package data\n\n")

    f.write("var TotalAgeWeights = " + str(totalAgeWeights) + "\n\n")
    f.write("var Ages = map[int]int{\n")
    min_weight = min(ages.values())
    for age, weight in ages.items():
        weight = max(1, int(weight/min_weight))
        f.write('\t' + str(age) + ': ' + str(weight) + ",\n")
    f.write("}\n")

with open("./faker/data/male_names.go", "w") as f:
    f.write("package data\n\n")
    f.write("var MaleNames = map[string]int{\n")
    min_weight = min(male_data.values())
    for name, weight in male_data.items():
        weight = max(1, int(weight/min_weight))
        name = name.capitalize()
        f.write('\t"' + name + '": ' + str(weight) + ",\n")
    f.write("}\n")

with open("./faker/data/female_names.go", "w") as f:
    f.write("package data\n\n")
    f.write("var FemaleNames = map[string]int{\n")
    min_weight = min(female_data.values())
    for name, weight in female_data.items():
        weight = max(1, int(weight/min_weight))
        name = name.capitalize()
        f.write('\t"' + name + '": ' + str(weight) + ",\n")
    f.write("}\n")

with open("./faker/data/last_names.go", "w") as f:
    f.write("package data\n\n")
    f.write("var LastNames = map[string]int{\n")
    min_weight = min(surnames.values())
    for name, weight in surnames.items():
        weight = max(1, int(weight/min_weight))
        name = name.capitalize()
        f.write('\t"' + name + '": ' + str(weight) + ",\n")
    f.write("}\n")
