![](Aspose.Words.4bd26d61-ac7b-4bd9-bb3f-2af61c60bffe.001.png)
# Selection Sort
O(n^2)
1) Najít nejmenší nesetřízený prvek
2) Vyměnit ho s prvním nesetřízeným prvkem
3) Opakovat pro zbývající prvky

|Word 10M|5m|0.000139%|
|-----------|------|------|
|Int 10M|5m|0.0002917%|
# Bubble Sort
O(n^2)
1) Projít pole, pokud je prvek větší než následující, vyměnit je, takhle ho probublat dokud to jde
2) Opakovat, dokud není pole seřazené

|Word 10M|5m|0.0001751%|
|-----------|------|------|
|Int 10M|5m|0.00038%|
# Insertion Sort
O(n^2)
1) Projít pole od druhého prvku
2) Vzít aktuální prvek a porovnat ho s předchozími, dokud nenajde správné místo
3) Vložit ho na správné místo
4) Opakovat pro všechny prvky

|Word 10M|5m|0.016508%|
|-----------|------|------|
|Int 10M|5m|0.0196492%|
# Heapsort
O(n log n)
1) Vytvořit haldu z pole
2) Opakovaně odebrat největší prvek z haldy a vložit ho na konec pole


|Word 10M|10.60s|
|-----------|------|
|Int 10M|15.46s|
# Merge Sort
O(n log n)
1) Rozdělit pole na dvě poloviny
2) Rekurzivně seřadit obě poloviny
3) Sloučit obě seřazené poloviny do jednoho pole

|Word 10M|7.92s|
|-----------|------|
|Int 10M|7.51s|
# Quick Sort (The words are worst case scenario, the problem is once the bar starts to fill it is essentially done, so there is a bit of a hack)
O(n log n) average, O(n^2) worst case
1) Vybrat pivot (např. poslední prvek)
2) Rozdělit pole na prvky menší a větší než pivot
3) Rekurzivně seřadit obě části
4) Sloučit seřazené části a pivot

|Word 10M|5m|0.0016566%|
|-----------|------|-----|
|Int 10M|5.33s|100%|
# Radix Sort
O(d (n + k)) kde d je nejdelší číslo, k je base
1) Seřadit čísla podle jednotlivých číslic od nejvíce v pravo
2) Použít Counting Sort pro každou číslici

Pls nepoužívat s floatama a číslama < 0 pokud pole zároveň neobsahuje string

|Word 10M|13.43225631s|
|-----------|------|
|Int 10M|560.298451ms|
