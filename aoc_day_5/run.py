def check_if_on_right_side(in_rules, in_page_number):
    for rule in in_rules:
        if rule[1] == in_page_number:
            return True

    return False

def remove_rules_for_page(in_rules, in_page_number):
    print("rules before: " + str(in_rules))
    deleted_rules = 0
    number_of_initial_rules = len(in_rules)
    i = 0
    while i < number_of_initial_rules:
        rule = in_rules[i-deleted_rules]
        if(rule[0] == in_page_number or rule[1] == in_page_number):
            # remove rule
            print("removing rule " + str(rule))
            in_rules.remove(rule)
            deleted_rules = deleted_rules + 1
        i = i + 1

    print("rules after: " + str(in_rules))
    return in_rules

def evaluate_update(update):
    len_of_update = len(update)
    middle_number = list(update.keys())[list(update.values()).index(len_of_update//2)]
    # print(middle_number)
    return int(middle_number)

def run():
    rules = []
    in_pth_rules = 'input/rules.txt'
    with open(in_pth_rules, 'r') as file:
        # Read each line in the file
        for line in file:
            rule_els = line.strip().split('|')
            rules.append([rule_els[0], rule_els[1]])

    # for rule in rules:
        # print("first el: " + rule[0] + " second el: " + rule[1])

    updates = []
    in_pth_updates = 'input/updates.txt'
    with open(in_pth_updates, 'r') as file:
        # Read each line in the file
        for line in file:
            # create dict
            tmp_dict = {}
            pages = line.strip().split(',')
            for index, page  in enumerate(pages):
                tmp_dict[page] = index
                # print("page " + str(page) + " at " + str(index))

            updates.append(tmp_dict)

    valid_updates = []
    formerly_invalid_updates = []
    for update in updates:
        # find all relevant rules
        rel_rules = []
        for rule in rules:
            if rule[0] in update and rule[1] in update:
                # add rule
                rel_rules.append(rule)

        all_rules_followed = True
        # loop through all relevant rules and check if rule is followed
        for rel_rule in rel_rules:
            # print("found relevant rule with first number " + str(rel_rule[0]) + " and second number " + str(rel_rule[1]))
            # get first index
            ind1 = update[rel_rule[0]]
            ind2 = update[rel_rule[1]]
            if ind1 > ind2:
                # rule has not been followed
                all_rules_followed = False
        
        if all_rules_followed:
            valid_updates.append(update)
        else: 
            # fix invalid update
            print("correcting update")
            update_size = len(update)
            i = 0
            corrected_update = {}
            while i < update_size:
                for page_number in update:
                    print("checking page number " + str(page_number))
                    if not check_if_on_right_side(rel_rules, page_number):
                        print("found new page " + str(page_number))
                        corrected_update[page_number] = i
                        # remove rules for the found page
                        rel_rules = remove_rules_for_page(rel_rules, page_number)
                        update.pop(page_number)
                        break
                i = i + 1
            formerly_invalid_updates.append(corrected_update)
    
    sum = 0
    # get all middle numbers of valid updates and add them to sum
    for valid_update in valid_updates:
        sum = sum + evaluate_update(valid_update)

    print("sum of valid updates: " + str(sum))

    # evaluate formerly invalid updates
    sum_inv = 0
    for update in formerly_invalid_updates:
        sum_inv = sum_inv + evaluate_update(update)
    
    print("sum of formerly invalid updates: " + str(sum_inv))


if __name__ == "__main__":
    run()