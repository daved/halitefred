#!/bin/bash

pdir="prep"
sdir="${pdir}/src"
hdir="${sdir}/github.com/daved/halitego"
mfile="MyBot.go"
zfile="sub_halitego.zip"

rm -rf ${pdir}
mkdir -p ${hdir}

cp ./cmd/gopherbot/main.go ${pdir}/${mfile}
cp -a ./vendor/* ${sdir}
cp -a ./geom ./internal/* ./ops ${hdir}

pushd ${pdir}

sed -i 's#/internal/#/#g' ${mfile}

zip -r ${zfile} ./*
mv ${zfile} ../

popd

rm -rf ${pdir}
