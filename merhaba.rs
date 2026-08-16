fn selam_ver(isim: &str) -> String {
    let mesaj = format!("Merhaba, {}!", isim);
    println!("{}", mesaj);
    mesaj
}

fn main() {
    selam_ver("Lonca");
}