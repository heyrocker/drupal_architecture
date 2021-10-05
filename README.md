# Drupal Architecture Utility

Drupal Architecture is a command-line utility which analyzes Drupal’s configuration files in order to generate comma-delimited files which describe various aspects of Drupal architecture. This can be very useful in order to get an overview of a site you are unfamiliar with. 

## Installation
Drupal Architecture is written in Go. If you are using a 64-bit version of OSX, then an executable called drupal_architecture is included in this repo. You should be able to download and run it as described below. For other platforms, you will need to install Go and either run the code directly or build an executable for your platform.

## Usage
```shell
drupal_architecture --read_dir --write_dir
```
- `--read_dir` "Full path to the directory containing your Drupal config files. Do not include trailing slash. Defaults to current directory.` 
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

### content_type_[name].csv
A listing of all the fields in a specific content type, as well as their properties.

- Label (Name of this field)
- Machine Name (Machine name of this field)
- Type (Field type of this field (text, entity reference, etc.))
- Description (Description of this field)
- Required (Is this field required?)
- Default Value (Is there a default value for this field?)
- Cardinality (Can there be more than one instance of this field?)
- Translatable (Is this field translatable?)

## paragraphs.csv
A listing of all the paragraphs in the Drupal installation. Contains the following fields

- Type (Name of this content type)
- Name (Machine name of this content type)
- Description (Description of this content type)

### paragraph_[name].csv
A listing of all the fields in a specific paragraph type, as well as their properties.

- Label (Name of this field)
- Machine Name (Machine name of this field)
- Type (Field type of this field (text, entity reference, etc.))
- Description (Description of this field)
- Required (Is this field required?)
- Default Value (Is there a default value for this field?)
- Cardinality (Can there be more than one instance of this field?)
- Translatable (Is this field translatable?)

## FAQ
### Why didn’t you write this as a Drupal module / Drush plugin?
The great thing about this utility as it stands now is that it doesn’t require a running Drupal site. You have a pile of config files and run it and you’re off. If I was going to write this as a Drupal module it would have most likely taken me longer to get a running development environment than it did to write the utility, and that is despite the fact that I had never written any Go before.

### Can you include [some other listings of things]
Probably! Send me your ideas, although as things stand, I have got what I need included.

### This code is terrible! Its like you’ve never written Go before!
I know! I haven’t! And yet it works! Amazing!

### Can I file a bug report? 
Sure! A pull request is even better! I have not done a ton of testing as I only have so many sets of config available to me. If you are experiencing a problem, please include your config with your bug report so I can test it.

### Can I write this as a Drupal module / Drush plugin?
Sure! Go crazy! That would be super useful for the community!