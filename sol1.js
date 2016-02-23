/*
 * Complete the function below.
 */
/*
For your reference:
LinkedListNode {
    var val;
    var next;
};
*/
function find(list, sublist) {
	var index = -1
	var i = list
	var j = sublist
	var match = true
	do {
		i = list
		j = sublist
		index += 1
		match = true
		do {
			if (i.val != j.val) {
				match = false
			}
			j = j.next
			i = i.next
			if (i == null && j != null) {
				match = false
			}
		} while (j != null && match == true)
		list = list.next
	} while (list != null && match == false)
	return index
}
