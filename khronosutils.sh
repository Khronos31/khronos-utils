yy(){
  if fc -ld >/dev/null 2>&1; then
    fc -e '(){cat "$1"&&echo>"$1"}' -1 2>/dev/null
  else
    fc -ln -1
  fi |
  sed -e '1s/^\s*//;:a;N;$!ba;s/\s*$//' |
  head -c -1 |
  clip
}
