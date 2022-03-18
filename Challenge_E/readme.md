Since we have 10 buckets the bucket weight is evenyl distributed among pens
ig. if a 100g buckets has 10 pens then each pen weights 10g assuming its not a faulty bucket
so the the algorithm would be same as the following
1- assign numbers to each bucket
2- for each bucket i, take i pens from the bucket
3- now we put the pens we took out of the  buckets on the scale, the perfect weight would be (1+2+3....+10)*10=550g
4- since we know that the faulty bucket weights 90g and contains 10 pens, so each pen weights 9gm, that means that the difference between a faulty pen and normal pen is 1g so we can deduce that the faulty bucket is simply the difference between the perfect weight and the actual weight.
for example 550-545 = 5 so the faulty bucket is bucket number 5
