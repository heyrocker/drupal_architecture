# Drupal Architecture Utility

Drupal Architecture is a command-line utility which analyzes Drupal’s configuration files in order to generate comma-delimited files which describe various aspects of Drupal architecture. This can be very useful in order to get an overview of a site you are unfamiliar with. 

## Installation
Drupal Architecture is written in Go. If you are using a 64-bit version of OSX, then an executable called drupal_architecture is included in this repo. You should be able to download and run it as described below. For other platforms, you will need to install Go and either run the code directly or build an executable for your platform.

## Usage
```shell
drupal_architecture --read_dir --write_dir
```
- `--read_dir ", "Full path to the directory containing your Drupal config files. Do not include trailing slash. Defaults to current directory.` 
- `--write_dir` "Full path to the directory where to write CSVs. Do not include trailing slash. Defaults to current directory"

## Output
When run, Drupal Architecture will write the following CSV files to write_dir

### content_types.csv
A listing of all the content types in the Drupal installation. Contains the following fields

- Type (Name of this content type)
- Name (Machine name of this content type)
- Description (Description of this content type)

### taxonomies.csv
A listing of all the taxonomy vocabularies in the Drupal installation. Contains the following fields

- Type (Name of this vocabulary)
- Name (Machine name of this vocabulary)
- Description (Description of this vocabulary)

### views.csv
A listing of all the views in the Drupal installation. Contains the following fields

- Label (Name of this view)
- Description (Description of this view)

### content_type_<name>.csv
A listing of all the fields in a specific content type, as well as their properties.

- Label (Name of this field)
- Machine Name (Machine name of this field)
- Type (Field type of this field (text, entity reference, etc.))
- Description (Description of this field)
- Required (Is this field required?)
- Cardinality (Can there be more than one instance of this field?)
- Translatable (Is this field translatable?)




