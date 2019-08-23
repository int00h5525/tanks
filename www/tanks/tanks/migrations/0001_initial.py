# Generated by Django 2.1.7 on 2019-08-23 23:00

import django.core.validators
from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Tank',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=64)),
                ('author', models.CharField(max_length=64)),
                ('color', models.CharField(max_length=7, validators=[django.core.validators.RegexValidator('#[0-9a-fA-F]{6}')])),
                ('program', models.CharField(max_length=16384)),
                ('s0r', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s0a', models.IntegerField(validators=[django.core.validators.MinValueValidator(-359), django.core.validators.MaxValueValidator(359)])),
                ('s0w', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s1r', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s1a', models.IntegerField(validators=[django.core.validators.MinValueValidator(-359), django.core.validators.MaxValueValidator(359)])),
                ('s1w', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s2r', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s2a', models.IntegerField(validators=[django.core.validators.MinValueValidator(-359), django.core.validators.MaxValueValidator(359)])),
                ('s2w', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s3r', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s3a', models.IntegerField(validators=[django.core.validators.MinValueValidator(-359), django.core.validators.MaxValueValidator(359)])),
                ('s3w', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s4r', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s4a', models.IntegerField(validators=[django.core.validators.MinValueValidator(-359), django.core.validators.MaxValueValidator(359)])),
                ('s4w', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s5r', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s5a', models.IntegerField(validators=[django.core.validators.MinValueValidator(-359), django.core.validators.MaxValueValidator(359)])),
                ('s5w', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s6r', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s6a', models.IntegerField(validators=[django.core.validators.MinValueValidator(-359), django.core.validators.MaxValueValidator(359)])),
                ('s6w', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s7r', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s7a', models.IntegerField(validators=[django.core.validators.MinValueValidator(-359), django.core.validators.MaxValueValidator(359)])),
                ('s7w', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s8r', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s8a', models.IntegerField(validators=[django.core.validators.MinValueValidator(-359), django.core.validators.MaxValueValidator(359)])),
                ('s8w', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s9r', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
                ('s9a', models.IntegerField(validators=[django.core.validators.MinValueValidator(-359), django.core.validators.MaxValueValidator(359)])),
                ('s9w', models.IntegerField(validators=[django.core.validators.MinValueValidator(0), django.core.validators.MaxValueValidator(100)])),
            ],
        ),
    ]