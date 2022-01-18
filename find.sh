#!/bin/bash
types=($(ls -d algorithm/*) exit)

chosen="${types[0]}"
select t in "${types[@]}"; do
  echo "You have chosen $t"
  chosen="$t"
  break
  [[ $t == exit ]] && exit 0
done

query="ls $chosen -1"
freqs=( all high medium low )
level=( all medium easy hard )
queryEnd=" |awk -F '_' \"{gsub ( \"[lc]\", \"\") ; print \$1}\""

select t in "${freqs[@]}"; do
  case $t in
    all )
      break
      ;;
    high )
      query=$query"| grep freqH"
      ;;
    medium )
      query=$query"| grep freqM" && break
      ;;
    low )
      query=$query"| grep freqL" && break
      ;;
    *)
      break
      ;;
  esac
done
select l in "${level[@]}"; do
  case $l in
    all)
      break
      ;;
    *)
      query=$query"| grep $l" && break
      ;;
  esac
done


findresult () {
  eval $query | awk -F "_" '{gsub ("[lc]", "") ; printf "%s%-50s%s%s%s\n","leetcode_", $0, "https://leetcode.com/problemset/all/?search=", $1, "&page=1"}'
}

findresult
exit 0
